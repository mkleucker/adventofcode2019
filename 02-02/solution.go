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

func handleAtPos(position int, numbers []int) (int, int) {
	opCode := numbers[position]
	value1 := numbers[numbers[position+1]]
	value2 := numbers[numbers[position+2]]
	target := numbers[position+3]

	if opCode == 1 {
		numbers[target] = value1 + value2
	} else if opCode == 2 {
		numbers[target] = value1 * value2
	} 

	return opCode, numbers[target]
}

func main() {  

	for input1 := 0; input1 < 100; input1++ {
		for input2 := 0; input2 < 100; input2++ {
			numbers := getData()

			numbers[1] = input1
			numbers[2] = input2
		
			for i := 0; i < len(numbers); i+= 4 {
				opCode, result := handleAtPos(i, numbers)
		
				if result == 19690720 {
					fmt.Printf("Input Values: %d / %d \n", input1, input2)
					fmt.Printf("Solution: %d \n", input1 * 100 + input2)
					
					input1 = 100
					input2 = 100
				}
				if opCode == 99 {
					i = len(numbers)
				}
			}
		
		}
	}
	
}
