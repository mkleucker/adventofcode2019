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
}
