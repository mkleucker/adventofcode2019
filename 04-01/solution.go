package main

import "fmt"

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
								if number >= start && (first == second || second == third || third == fourth || fourth == fifth || fifth == sixth) {
									count++
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
