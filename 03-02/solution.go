package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
func getPoints(startX int, startY int, moveX int, moveY int) []string {
	visitedPoints := []string{}

	factorX := 1
	if moveX < 0 {
		factorX = -1
	}
	factorY := 1
	if moveY < 0 {
		factorY = -1
	}
	for x := 0; x <= (startX+moveX*factorX)-startX; x++ {
		for y := 0; y <= (startY+moveY*factorY)-startY; y++ {
			visitedPoints = append(visitedPoints, strconv.Itoa(startX+x*factorX)+"|"+strconv.Itoa(startY+y*factorY))
		}
	}

	return visitedPoints
}

func getWirePath(wire []string) []string {
	path := []string{}

	currentPosX := 0
	currentPosY := 0

	for _, instruction := range wire {
		moveX, moveY := handleInstruction(instruction)

		path = append(path, getPoints(currentPosX, currentPosY, moveX, moveY)...)
		// remove the last item as it will be the first one of the next segment and then counted twice
		path = path[:len(path)-1]
		currentPosX = currentPosX + moveX
		currentPosY = currentPosY + moveY
	}

	return path
}

func getIntersectingPoints(path1 []string, path2 []string) map[string]int {
	crossings := make(map[string]int)
	for i, e := range path1 {
		if e == "0|0" {
			continue
		}

		for j, e2 := range path2 {
			if e2 == "0|0" {
				continue
			}
			if e == e2 {
				crossings[e] = i + j
			}
		}
	}
	return crossings
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
	for _, distance := range intersections {
		shortestDistance = min(shortestDistance, distance)
	}

	fmt.Printf("Solution: %d \n", shortestDistance)
}
