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

func calculateImage(layers [][]int) []int {
	image := layers[0]
	for _, layer := range layers[1:] {
		for pixelIndex, pixel := range layer {
			if image[pixelIndex] == 2 {
				image[pixelIndex] = pixel
			}
		}
	}

	return image
}

func drawImage(pixels []int, width int, height int) {
	for i, pixel := range pixels {
		if i%(width) == 0 {
			fmt.Printf("\n")
		}
		if pixel == 0 {
			fmt.Printf("\u2591\u2591")
		} else if pixel == 1 {
			fmt.Printf("  ")
		} else {
			fmt.Printf("  ")
		}
	}
	fmt.Printf("\n")
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

	width := 25
	height := 6

	layers := extractLayers(imageData, width, height)

	image := calculateImage(layers)

	drawImage(image, width, height)
}
