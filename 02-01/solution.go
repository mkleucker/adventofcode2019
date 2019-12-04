package main

import (  
    "fmt"
		"io/ioutil"
		"strings"
		"strconv"
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

func handleAtPos(position int, numbers []int) int {
	opCode := numbers[position]
	value1 := numbers[numbers[position+1]]
	value2 := numbers[numbers[position+2]]
	target := numbers[position+3]

	if opCode == 1 {
		numbers[target] = value1 + value2
	} else if opCode == 2 {
		numbers[target] = value1 * value2
	} 

	return opCode
}

func main() {  
	numbers := getData()

	numbers[1] = 12
	numbers[2] = 2

	for i := 0; i < len(numbers); i+= 4 {
		opCode := handleAtPos(i, numbers)

		if opCode == 99 {
			i = len(numbers)
		}
	}

	fmt.Printf("Solution: %d", numbers[0])
}
