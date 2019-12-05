package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	start := 245182
	end := 790572

	firstStart := start / 100000
	firstEnd := end / 100000

	count := 0

	for first := firstStart; first <= firstEnd; first++ {
		for second := first; second <= 9; second++ {
			for third := second; third <= 9; third++ {
				for fourth := third; fourth <= 9; fourth++ {
					for fifth := fourth; fifth <= 9; fifth++ {
						for sixth := fifth; sixth <= 9; sixth++ {
							number := first*100000 + second*10000 + third*1000 + fourth*100 + fifth*10 + sixth
							if number <= end {
								if number >= start {
									numberAsString := strconv.Itoa(number)

									for i := 1; i <= 9; i++ {
										double := strconv.Itoa(i * 11)
										triple := strconv.Itoa(i * 111)

										if strings.Contains(numberAsString, double) && !strings.Contains(numberAsString, triple) {
											count++
											break
										}
									}
								}
							} else {
								break
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("Solution: %d \n", count)

}
