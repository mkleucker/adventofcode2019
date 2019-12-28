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

func (amp *Amp) add(paramMode1 int, paramMode2 int, paramMode3 int) {
	val1 := amp.getValue(paramMode1, amp.CurrentPos+1)
	val2 := amp.getValue(paramMode2, amp.CurrentPos+2)

	amp.setValue(paramMode3, amp.CurrentPos+3, (val1 + val2))
	amp.CurrentPos += 4
}

func (amp *Amp) multiply(paramMode1 int, paramMode2 int, paramMode3 int) {
	val1 := amp.getValue(paramMode1, amp.CurrentPos+1)
	val2 := amp.getValue(paramMode2, amp.CurrentPos+2)

	amp.setValue(paramMode3, amp.CurrentPos+3, (val1 * val2))
	amp.CurrentPos += 4
}

func (amp *Amp) jumpIfTrue(paramMode1 int, paramMode2 int) (bool, int) {
	val1 := amp.getValue(paramMode1, amp.CurrentPos+1)

	if val1 == 0 {
		return false, 0
	}

	val2 := amp.getValue(paramMode2, amp.CurrentPos+2)

	return true, val2
}

func (amp *Amp) jumpIfFalse(paramMode1 int, paramMode2 int) (bool, int) {
	val1 := amp.getValue(paramMode1, amp.CurrentPos+1)

	if val1 != 0 {
		return false, 0
	}

	val2 := amp.getValue(paramMode2, amp.CurrentPos+2)

	return true, val2
}

func (amp *Amp) lessThan(paramMode1 int, paramMode2 int, paramMode3 int) {
	val1 := amp.getValue(paramMode1, amp.CurrentPos+1)
	val2 := amp.getValue(paramMode2, amp.CurrentPos+2)

	newValue := 0
	if val1 < val2 {
		newValue = 1
	}
	amp.setValue(paramMode3, amp.CurrentPos+3, newValue)
	amp.CurrentPos += 4
}

func (amp *Amp) equals(paramMode1 int, paramMode2 int, paramMode3 int) {
	val1 := amp.getValue(paramMode1, amp.CurrentPos+1)
	val2 := amp.getValue(paramMode2, amp.CurrentPos+2)

	newValue := 0
	if val1 == val2 {
		newValue = 1
	}

	amp.setValue(paramMode3, amp.CurrentPos+3, newValue)
	amp.CurrentPos += 4
}

func (amp *Amp) getValue(mode int, position int) int {
	if mode == 1 {
		return amp.Numbers[position]
	} else if mode == 2 {
		return amp.Numbers[amp.RelativeBase+amp.Numbers[position]]
	}
	return amp.Numbers[amp.Numbers[position]]
}

func (amp *Amp) setValue(mode int, position int, value int) {
	if mode == 1 {
		amp.Numbers[position] = value
	} else if mode == 2 {
		amp.Numbers[amp.RelativeBase+amp.Numbers[position]] = value
	} else {
		amp.Numbers[amp.Numbers[position]] = value
	}
}

func (amp *Amp) adjustRelativeBase(paramMode int) {
	baseAdjust := amp.getValue(paramMode, amp.CurrentPos+1)
	amp.RelativeBase += baseAdjust
	amp.CurrentPos += 2
}

type Amp struct {
	CurrentPos   int
	Numbers      []int
	Done         bool
	LastOutput   int
	RelativeBase int
}

func (amp *Amp) run(input int) {
	for {
		instruction := amp.Numbers[amp.CurrentPos]
		opCode, paramMode1, paramMode2, paramMode3 := interpretInstruction(instruction)

		if opCode == 99 {
			amp.Done = true
			return
		} else if opCode == 1 {
			amp.add(paramMode1, paramMode2, paramMode3)
		} else if opCode == 2 {
			amp.multiply(paramMode1, paramMode2, paramMode3)
		} else if opCode == 3 {
			if paramMode1 == 1 {
				amp.Numbers[amp.CurrentPos+1] = input
			} else if paramMode1 == 2 {
				amp.Numbers[amp.RelativeBase+amp.Numbers[amp.CurrentPos+1]] = input
			} else {
				amp.Numbers[amp.Numbers[amp.CurrentPos+1]] = input
			}
			amp.CurrentPos += 2
		} else if opCode == 4 {
			amp.LastOutput = amp.getValue(paramMode1, amp.CurrentPos+1)
			amp.CurrentPos += 2
			return
		} else if opCode == 5 {
			jump, newPosition := amp.jumpIfTrue(paramMode1, paramMode2)
			if jump {
				amp.CurrentPos = newPosition
			} else {
				amp.CurrentPos += 3
			}
		} else if opCode == 6 {
			jump, newPosition := amp.jumpIfFalse(paramMode1, paramMode2)
			if jump {
				amp.CurrentPos = newPosition
			} else {
				amp.CurrentPos += 3
			}
		} else if opCode == 7 {
			amp.lessThan(paramMode1, paramMode2, paramMode3)
		} else if opCode == 8 {
			amp.equals(paramMode1, paramMode2, paramMode3)
		} else if opCode == 9 {
			amp.adjustRelativeBase(paramMode1)
		}
	}
}

func initRun() int {

	amp := &Amp{
		Numbers: getData(),
	}
	amp.run(2)

	return amp.LastOutput
}

func main() {

	output := initRun()

	fmt.Printf("Program Output: %d\n", output)

}
