package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getData() []int {
	file, _ := ioutil.ReadFile("data.txt")
	data := string(file)
	stringData := strings.Split(data, ",")
	numbers := []int{}

	for _, i := range stringData {
		j, _ := strconv.Atoi(i)
		numbers = append(numbers, j)
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

func add(numbers []int, position int, val1Mode int, val2Mode int) []int {
	val1 := 0
	if val1Mode == 1 {
		val1 = numbers[position+1]
	} else {
		val1 = numbers[numbers[position+1]]
	}
	val2 := 0
	if val2Mode == 1 {
		val2 = numbers[position+2]
	} else {
		val2 = numbers[numbers[position+2]]
	}

	numbers[numbers[position+3]] = val1 + val2

	return numbers
}

func multiply(numbers []int, position int, val1Mode int, val2Mode int) []int {
	val1 := 0
	if val1Mode == 1 {
		val1 = numbers[position+1]
	} else {
		val1 = numbers[numbers[position+1]]
	}
	val2 := 0
	if val2Mode == 1 {
		val2 = numbers[position+2]
	} else {
		val2 = numbers[numbers[position+2]]
	}

	numbers[numbers[position+3]] = val1 * val2
	return numbers
}

func main() {

	numbers := getData()

	instructionLength := 0
	input := 1

	for i := 0; i < len(numbers); i += instructionLength {
		instruction := numbers[i]
		opCode, paramMode1, paramMode2, _ := interpretInstruction(instruction)

		if opCode == 99 {
			fmt.Printf("Done \n")
			os.Exit(0)
		} else if opCode == 1 {
			numbers = add(numbers, i, paramMode1, paramMode2)
		} else if opCode == 2 {
			numbers = multiply(numbers, i, paramMode1, paramMode2)
		} else if opCode == 3 {
			numbers[i+1] = input
		} else if opCode == 4 {
			if paramMode1 == 1 {
				fmt.Printf("Output: %d \n", numbers[i+1])
			} else {
				fmt.Printf("Output: %d \n", numbers[numbers[i+1]])
			}
		}
		if opCode > 2 {
			instructionLength = 2
		} else {
			instructionLength = 4
		}
	}

}
