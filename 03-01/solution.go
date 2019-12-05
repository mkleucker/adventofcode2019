package main

import (  
    "fmt"
		"os"
		"bufio"
		"strings"
		"strconv"
)


func getData() [][]string {
	wires := [][]string{}

	file, _ := os.Open("data.txt")
		
	scanner := bufio.NewScanner(file)

	// iterate over line
	for scanner.Scan() {
		line := scanner.Text()
		wire := strings.Split(line, ",")
		wires = append(wires, wire)
	}

	return wires
}

func handleInstruction(instruction string) (int, int) {
	startX := 0
	startY := 0

	direction := instruction[:1]
	distance, _ := strconv.Atoi(instruction[1:])

	if direction == "U" {
		return startX, (startY + distance)
	} else if direction == "D" {
		return startX, (startY - distance)
	} else if direction == "L" {
		return (startX - distance), startY
	} else if direction == "R" {
		return (startX + distance), startY
	}
	return startX, startY
}	

func min(x, y int) int {
	if x < y {
			return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
			return x
	}
	return y
}
func getPoints(startX int, startY int, endX int, endY int) []string {
	visitedPoints := []string{}

	for x := min(startX, endX); x <= max(startX, endX); x++ {
		for y := min(startY, endY); y <= max(startY, endY); y++ {
			visitedPoints = append(visitedPoints, strconv.Itoa(x)+"|"+strconv.Itoa(y))
		}		
	}

	return visitedPoints
}

func getWirePath(wire []string) []string {
	path :=  []string {}

	currentPosX := 0
	currentPosY := 0

	for _, instruction := range wire {
		moveX, moveY := handleInstruction(instruction)
		newX := currentPosX + moveX
		newY := currentPosY + moveY

		path = append(path, getPoints(currentPosX, currentPosY, newX, newY)...)
		currentPosX = newX
		currentPosY = newY
	}

	return path
}

func getIntersectingPoints(path1 []string, path2 []string) []string {
	crossings := []string {}
	for _, e := range path1 {
		if e == "0|0" {
			continue
		}
		for _, e2 := range path2 {
			if e == e2 {
				crossings = append(crossings, e)
				continue
			}
		}
	}
	return crossings
}

func getDistance(point string) int {
	pointData := strings.Split(point, "|")
	x,_ := strconv.Atoi(pointData[0])
	y,_ := strconv.Atoi(pointData[1])

	return max(x, -1 *x) + max(y, -1*y)
}


func main() {  
	wires := getData()

	paths := []([]string){}
	for _, wire := range wires {
		path := getWirePath(wire)
		paths = append(paths, path)
	}

	intersections := getIntersectingPoints(paths[0], paths[1])

	shortestDistance := 99999999
	for _, intersection := range intersections {
		shortestDistance = min(shortestDistance, getDistance(intersection))
	}

	fmt.Printf("Solution: %d \n", shortestDistance)
}
