package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func getData() [][]bool {
	file, _ := os.Open("data.txt")

	scanner := bufio.NewScanner(file)

	asteroids := [][]bool{}
	// iterate over line
	for scanner.Scan() {
		lineText := scanner.Text()
		lineMap := []bool{}
		for i := 0; i < len(lineText); i++ {
			if string(lineText[i]) == "#" {
				lineMap = append(lineMap, true)
			} else {
				lineMap = append(lineMap, false)
			}
		}
		asteroids = append(asteroids, lineMap)
	}

	return asteroids
}

func getAngle(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Atan2(float64(y2-y1), float64(x2-x1))
}

type Asteroid struct {
	X int
	Y int
}

func main() {

	asteroidMap := getData()
	height := len(asteroidMap)
	width := len(asteroidMap[0])

	asteroids := []Asteroid{}

	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			if asteroidMap[x][y] == true {
				asteroid := Asteroid{
					X: x,
					Y: y,
				}
				asteroids = append(asteroids, asteroid)
			}
		}
	}

	maxVisible := 0
	for _, monitoringAsteroid := range asteroids {
		visibleAngles := []float64{}
		for _, outerAsteroid := range asteroids {
			if outerAsteroid.X == monitoringAsteroid.X && outerAsteroid.Y == monitoringAsteroid.Y {
				continue
			}
			angle := getAngle(monitoringAsteroid.X, monitoringAsteroid.Y, outerAsteroid.X, outerAsteroid.Y)
			alreadyKnown := false
			for _, knownAngle := range visibleAngles {
				if knownAngle == angle {
					alreadyKnown = true
				}
			}
			if !alreadyKnown {
				visibleAngles = append(visibleAngles, angle)
			}
		}
		if len(visibleAngles) > maxVisible {
			maxVisible = len(visibleAngles)
		}
	}

	fmt.Printf("Visible %d", maxVisible)
}
