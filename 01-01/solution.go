package main

import (  
    "fmt"
		"os"
		"bufio"
		"strconv"
)

func calcFuel(mass int) int {
	requiredFuel := mass / 3 - 2

	if requiredFuel > 0 {
		return requiredFuel
	}
	return 0
}

func main() {  
	file, _ := os.Open("data.txt")
		
	scanner := bufio.NewScanner(file)

	totalFuel := 0
	// iterate over line
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		totalFuel += calcFuel(number)
	}

	fmt.Printf("Solution: %d", totalFuel)
}
