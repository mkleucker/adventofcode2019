package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func getData() []int {
	file, _ := ioutil.ReadFile("data.txt")
	data := string(file)
	numbers := []int{}
	for i := 0; i < len(data); i++ {
		j, _ := strconv.Atoi(string(data[i]))
		numbers = append(numbers, j)
	}

	return numbers
}

func extractLayers(imageData []int, width int, height int) [][]int {
	layers := [][]int{}
	numPixels := width * height
	numLayers := len(imageData) / numPixels

	for i := 0; i < numLayers; i++ {
		start := i * numPixels
		end := start + numPixels

		layer := imageData[start:end]
		layers = append(layers, layer)
	}
	return layers
}

func getLayerWithFewestZeroes(layers [][]int) []int {
	lowestZeroLayer := []int{}
	lowestZeroCount := len(layers[0])
	for _, layer := range layers {
		zeroes := 0
		for _, pixel := range layer {
			if pixel == 0 {
				zeroes++
			}
		}
		if zeroes < lowestZeroCount {
			lowestZeroCount = zeroes
			lowestZeroLayer = layer
		}
	}

	return lowestZeroLayer
}

func calculateChecksum(layer []int) int {
	ones := 0
	twos := 0

	for _, pixel := range layer {
		if pixel == 1 {
			ones++
		}
		if pixel == 2 {
			twos++
		}
	}

	return ones * twos
}
func main() {

	imageData := getData()

	layers := extractLayers(imageData, 25, 6)

	lowestZeroLayer := getLayerWithFewestZeroes(layers)

	checksum := calculateChecksum(lowestZeroLayer)

	fmt.Printf("Checksum %v\n", checksum)

}
