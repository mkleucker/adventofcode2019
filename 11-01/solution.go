package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func getData() []int {
	file, _ := ioutil.ReadFile("data.txt")
	data := string(file)
	stringData := strings.Split(data, ",")
	numbers := make([]int, math.MaxInt16)

	for index, i := range stringData {
		j, _ := strconv.Atoi(i)
		numbers[index] = j
	}

	return numbers
}

func interpretInstruction(instruction int) (int, int, int, int) {
	opCode := instruction % 100
	paramMode1 := 0
	paramMode2 := 0
	paramMode3 := 0

	if instruction > 100 {
		paramMode1 = (instruction%1000 - opCode) / 100
	}

	if instruction > 1000 {
		paramMode2 = (instruction%10000 - paramMode1*100 - opCode) / 1000
	}

	if instruction > 10000 {
		paramMode3 = (instruction%100000 - paramMode2*1000 - paramMode1*100 - opCode) / 10000
	}

	return opCode, paramMode1, paramMode2, paramMode3
}

func (computer *Computer) add(paramMode1 int, paramMode2 int, paramMode3 int) {
	val1 := computer.getValue(paramMode1, computer.CurrentPos+1)
	val2 := computer.getValue(paramMode2, computer.CurrentPos+2)

	computer.setValue(paramMode3, computer.CurrentPos+3, (val1 + val2))
	computer.CurrentPos += 4
}

func (computer *Computer) multiply(paramMode1 int, paramMode2 int, paramMode3 int) {
	val1 := computer.getValue(paramMode1, computer.CurrentPos+1)
	val2 := computer.getValue(paramMode2, computer.CurrentPos+2)

	computer.setValue(paramMode3, computer.CurrentPos+3, (val1 * val2))
	computer.CurrentPos += 4
}

func (computer *Computer) jumpIfTrue(paramMode1 int, paramMode2 int) (bool, int) {
	val1 := computer.getValue(paramMode1, computer.CurrentPos+1)

	if val1 == 0 {
		return false, 0
	}

	val2 := computer.getValue(paramMode2, computer.CurrentPos+2)

	return true, val2
}

func (computer *Computer) jumpIfFalse(paramMode1 int, paramMode2 int) (bool, int) {
	val1 := computer.getValue(paramMode1, computer.CurrentPos+1)

	if val1 != 0 {
		return false, 0
	}

	val2 := computer.getValue(paramMode2, computer.CurrentPos+2)

	return true, val2
}

func (computer *Computer) lessThan(paramMode1 int, paramMode2 int, paramMode3 int) {
	val1 := computer.getValue(paramMode1, computer.CurrentPos+1)
	val2 := computer.getValue(paramMode2, computer.CurrentPos+2)

	newValue := 0
	if val1 < val2 {
		newValue = 1
	}
	computer.setValue(paramMode3, computer.CurrentPos+3, newValue)
	computer.CurrentPos += 4
}

func (computer *Computer) equals(paramMode1 int, paramMode2 int, paramMode3 int) {
	val1 := computer.getValue(paramMode1, computer.CurrentPos+1)
	val2 := computer.getValue(paramMode2, computer.CurrentPos+2)

	newValue := 0
	if val1 == val2 {
		newValue = 1
	}

	computer.setValue(paramMode3, computer.CurrentPos+3, newValue)
	computer.CurrentPos += 4
}

func (computer *Computer) getValue(mode int, position int) int {
	if mode == 1 {
		return computer.Numbers[position]
	} else if mode == 2 {
		return computer.Numbers[computer.RelativeBase+computer.Numbers[position]]
	}
	return computer.Numbers[computer.Numbers[position]]
}

func (computer *Computer) setValue(mode int, position int, value int) {
	if mode == 1 {
		computer.Numbers[position] = value
	} else if mode == 2 {
		computer.Numbers[computer.RelativeBase+computer.Numbers[position]] = value
	} else {
		computer.Numbers[computer.Numbers[position]] = value
	}
}

func (computer *Computer) adjustRelativeBase(paramMode int) {
	baseAdjust := computer.getValue(paramMode, computer.CurrentPos+1)
	computer.RelativeBase += baseAdjust
	computer.CurrentPos += 2
}

type Computer struct {
	CurrentPos   int
	Numbers      []int
	Done         bool
	LastOutput   int
	RelativeBase int
}

func (computer *Computer) run(input int) {
	for {
		instruction := computer.Numbers[computer.CurrentPos]
		opCode, paramMode1, paramMode2, paramMode3 := interpretInstruction(instruction)

		if opCode == 99 {
			computer.Done = true
			return
		} else if opCode == 1 {
			computer.add(paramMode1, paramMode2, paramMode3)
		} else if opCode == 2 {
			computer.multiply(paramMode1, paramMode2, paramMode3)
		} else if opCode == 3 {
			if paramMode1 == 1 {
				computer.Numbers[computer.CurrentPos+1] = input
			} else if paramMode1 == 2 {
				computer.Numbers[computer.RelativeBase+computer.Numbers[computer.CurrentPos+1]] = input
			} else {
				computer.Numbers[computer.Numbers[computer.CurrentPos+1]] = input
			}
			computer.CurrentPos += 2
		} else if opCode == 4 {
			computer.LastOutput = computer.getValue(paramMode1, computer.CurrentPos+1)
			computer.CurrentPos += 2
			return
		} else if opCode == 5 {
			jump, newPosition := computer.jumpIfTrue(paramMode1, paramMode2)
			if jump {
				computer.CurrentPos = newPosition
			} else {
				computer.CurrentPos += 3
			}
		} else if opCode == 6 {
			jump, newPosition := computer.jumpIfFalse(paramMode1, paramMode2)
			if jump {
				computer.CurrentPos = newPosition
			} else {
				computer.CurrentPos += 3
			}
		} else if opCode == 7 {
			computer.lessThan(paramMode1, paramMode2, paramMode3)
		} else if opCode == 8 {
			computer.equals(paramMode1, paramMode2, paramMode3)
		} else if opCode == 9 {
			computer.adjustRelativeBase(paramMode1)
		}
	}
}

/**
0: "Up",
1: "Right",
2: "Down",
3: "Left",
*/

func adjust(direction int, adjustment int) int {
	if adjustment == 0 {
		adjustment = -1
	}
	if direction+adjustment < 0 {
		return 3
	}
	if direction+adjustment > 3 {
		return 0
	}
	return direction + adjustment
}

func calculateLocation(x, y, direction int) (int, int) {
	switch direction {
	case 0:
		return x, y - 1
	case 1:
		return x + 1, y
	case 2:
		return x, y + 1
	case 3:
		return x - 1, y
	}
	return x, y
}

func getInputFromGrid(grid map[int]map[int]int, x int, y int) int {
	if horizontal, ok := grid[x]; ok {
		if value, ok2 := horizontal[y]; ok2 {
			return value
		}
	}
	return 0
}

func initRun() int {

	computer := &Computer{
		Numbers: getData(),
	}

	currentDirection := 0

	grid := make(map[int]map[int]int)

	currentLocationX, currentLocationY := 0, 0

	for !computer.Done {
		input := getInputFromGrid(grid, currentLocationX, currentLocationY)
		computer.run(input)
		paint := computer.LastOutput
		computer.run(input)

		if _, ok := grid[currentLocationX]; !ok {
			grid[currentLocationX] = make(map[int]int)
		}

		grid[currentLocationX][currentLocationY] = paint

		currentDirection = adjust(currentDirection, computer.LastOutput)
		currentLocationX, currentLocationY = calculateLocation(currentLocationX, currentLocationY, currentDirection)
	}

	paintedPanels := 0

	for _, row := range grid {
		paintedPanels += len(row)
	}

	return paintedPanels
}

func main() {

	output := initRun()

	fmt.Printf("Program Output: %d\n", output)

}
