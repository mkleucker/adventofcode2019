package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addOrbit(orbit string) (string, string) {
	data := strings.Split(orbit, ")")
	return data[0], data[1]
}

func findStarting(orbits map[string][]string) string {
	for key := range orbits {
		found := false
		for _, value2 := range orbits {
			for _, orbit := range value2 {
				if key == orbit {
					found = true
				}
			}
		}
		if !found {
			return key
		}
	}
	return "meh"
}

func meh(orbits map[string][]string, starting string, level int) int {
	connections := orbits[starting]
	countConnections := len(connections)

	currentCount := level * countConnections

	for _, connection := range connections {
		currentCount += meh(orbits, connection, level+1)
	}

	return currentCount
}

func trace(orbits map[string][]string, targetOrbit string, startOrbit string, foundPaths *[]string, currentPath []string) {
	localFoundPaths := *foundPaths
	for _, orbit := range orbits[startOrbit] {
		if orbit == targetOrbit {
			currentPath = append(currentPath, orbit)
			*foundPaths = append(localFoundPaths, currentPath...)
		} else {
			currentPath = append(currentPath, orbit)
			trace(orbits, targetOrbit, orbit, foundPaths, currentPath)
			currentPath = currentPath[:len(currentPath)-1]
		}
	}
}

func lastOverlap(pathA []string, pathB []string) string {
	overlap := []string{}
	for _, elA := range pathA {
		for _, elB := range pathB {
			if elA == elB {
				overlap = append(overlap, elA)
			}
		}
	}

	return overlap[len(overlap)-1]
}

func getPathPartial(path []string, start string, end string) []string {
	for i, el := range path {
		if el == start {
			return path[i:]
		}
	}
	return []string{}
}

func main() {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)

	orbits := make(map[string][]string)
	// iterate over line
	for scanner.Scan() {
		line := scanner.Text()
		orbitFrom, orbitTo := addOrbit(line)
		if _, ok := orbits[orbitFrom]; !ok {
			orbits[orbitFrom] = []string{}
		}
		orbits[orbitFrom] = append(orbits[orbitFrom], orbitTo)
	}

	startingOrbit := findStarting(orbits)
	fmt.Printf("starting %s %d \n", startingOrbit, meh(orbits, startingOrbit, 1))

	pathA := []string{}
	trace(orbits, "YOU", startingOrbit, &pathA, []string{"COM"})
	pathB := []string{}
	trace(orbits, "SAN", startingOrbit, &pathB, []string{"COM"})

	roadFork := lastOverlap(pathA, pathB)

	partialA := getPathPartial(pathA, roadFork, "YOU")
	partialB := getPathPartial(pathB, roadFork, "SAN")

	fmt.Printf("Total Distance %d \n", (len(partialA) + len(partialB) - 4))

}
