package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func getAngle(x1, y1, x2, y2 int) float64 {
	angle := math.Atan2(float64(y2-y1), float64(x2-x1))
	if angle < 0 {
		angle = angle + (2 * math.Pi)
	}
	return angle * (180 / math.Pi)
}

func getDistance(x1, y1, x2, y2 int) float64 {
	a := math.Pow(float64(x2-x1), 2)
	b := math.Pow(float64(y2-y1), 2)
	return math.Sqrt(a + b)
}

type byAngle []float64

func (f byAngle) Len() int {
	return len(f)
}
func (f byAngle) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f byAngle) Less(i, j int) bool {
	// match the inverted coordinate system
	newI := f[i] - 270
	if newI < 0 {
		newI = newI + 360
	}
	newJ := f[j] - 270
	if newJ < 0 {
		newJ = newJ + 360
	}

	return newI < newJ
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
					X: y,
					Y: x,
				}
				asteroids = append(asteroids, asteroid)
			}
		}
	}

	maxVisible := 0
	var astroidDirections map[float64]map[float64]Asteroid

	for _, monitoringAsteroid := range asteroids {
		outerAsteroids := make(map[float64]map[float64]Asteroid)
		for _, outerAsteroid := range asteroids {
			if outerAsteroid.X == monitoringAsteroid.X && outerAsteroid.Y == monitoringAsteroid.Y {
				continue
			}
			angle := getAngle(monitoringAsteroid.X, monitoringAsteroid.Y, outerAsteroid.X, outerAsteroid.Y)
			distance := getDistance(monitoringAsteroid.X, monitoringAsteroid.Y, outerAsteroid.X, outerAsteroid.Y)

			if _, ok := outerAsteroids[angle]; !ok {
				outerAsteroids[angle] = make(map[float64]Asteroid)
			}
			outerAsteroids[angle][distance] = outerAsteroid
		}
		if len(outerAsteroids) > maxVisible {
			maxVisible = len(outerAsteroids)
			astroidDirections = outerAsteroids
		}
	}

	destroyCounter := 1

	destroyed := make(map[int]Asteroid)

	for {

		directions := []float64{}
		for direction := range astroidDirections {
			directions = append(directions, direction)
		}
		sort.Sort(byAngle(directions))

		for _, direction := range directions {
			asteroids := astroidDirections[direction]
			distances := []float64{}
			for distance := range asteroids {
				distances = append(distances, distance)
			}
			sort.Float64s(distances)

			shortestDistance := distances[0]

			destroyed[destroyCounter] = astroidDirections[direction][shortestDistance]

			delete(astroidDirections[direction], shortestDistance)
			if len(astroidDirections[direction]) == 0 {
				delete(astroidDirections, direction)
			}
			destroyCounter++
		}

		if len(astroidDirections) == 0 {
			break
		}
	}

	fmt.Printf("Destroy #200: %v %v\n", destroyed[200].X, destroyed[200].Y)
	fmt.Printf("Solution: %v\n", destroyed[200].X*100+destroyed[200].Y)
}
