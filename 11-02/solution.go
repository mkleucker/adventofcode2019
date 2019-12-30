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

type Location struct {
	X int
	Y int
}

func calculateLocation(loc Location, direction int) Location {
	switch direction {
	case 0:
		return Location{
			X: loc.X,
			Y: loc.Y - 1,
		}
	case 1:
		return Location{
			X: loc.X + 1,
			Y: loc.Y,
		}
	case 2:
		return Location{
			X: loc.X,
			Y: loc.Y + 1,
		}
	case 3:
		return Location{
			X: loc.X - 1,
			Y: loc.Y,
		}
	}
	return Location{
		X: loc.X,
		Y: loc.Y,
	}
}

func getInputFromGrid(grid map[Location]int, x int, y int) int {
	if pixel, ok := grid[Location{X: x, Y: y}]; ok {
		return pixel
	}
	return 0
}

func initRun() map[Location]int {

	computer := &Computer{
		Numbers: getData(),
	}

	currentDirection := 0

	grid := make(map[Location]int)

	grid[Location{X: 0, Y: 0}] = 1

	currentLocation := Location{X: 0, Y: 0}

	for !computer.Done {
		input := getInputFromGrid(grid, currentLocation.X, currentLocation.Y)
		computer.run(input)
		paint := computer.LastOutput
		computer.run(input)

		grid[currentLocation] = paint

		currentDirection = adjust(currentDirection, computer.LastOutput)
		currentLocation = calculateLocation(currentLocation, currentDirection)
	}

	return grid
}

func drawImage(grid map[Location]int, startX int, endX int, startY int, endY int) {
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {

			pixelColor := getInputFromGrid(grid, x, y)
			if pixelColor == 0 {
				fmt.Printf("\u2591")
			} else if pixelColor == 1 {
				fmt.Printf(" ")
			} else {
				fmt.Printf(" ")
			}

			if x == endX {
				fmt.Printf("\n")
			}
		}
	}

	fmt.Printf("\n")
}

func convertGridToImage(grid map[Location]int) {
	widthLow := 0
	widthHigh := 0
	heightLow := 0
	heightHigh := 0

	for location, _ := range grid {
		if location.X > widthHigh {
			widthHigh = location.X
		}
		if location.X < widthLow {
			widthLow = location.X
		}
		if location.Y > heightHigh {
			heightHigh = location.Y
		}
		if location.Y < heightLow {
			heightLow = location.Y
		}
	}

	drawImage(grid, widthLow, widthHigh, heightLow, heightHigh)
}

func main() {

	grid := initRun()

	convertGridToImage(grid)
	// fmt.Printf("Program Output: %d\n", output)

}
