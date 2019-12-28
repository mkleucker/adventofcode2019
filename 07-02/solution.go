package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getData() []int {
	file, _ := ioutil.ReadFile("data.txt")
	data := string(file)
	stringData := strings.Split(data, ",")
	numbers := []int{}

	for _, i := range stringData {
		j, _ := strconv.Atoi(i)
		numbers = append(numbers, j)
	}

	return numbers
}

func interpretInstruction(instruction int) (int, int, int, int) {
	opCode := instruction % 100
	paramMode1 := 0
	paramMode2 := 0
	paramMode3 := 0

	if instruction > 100 {
		paramMode1 = (instruction%1000 - opCode) / 100
	}

	if instruction > 1000 {
		paramMode2 = (instruction%10000 - paramMode1*100 - opCode) / 1000
	}

	if instruction > 10000 {
		paramMode3 = (instruction%100000 - paramMode2*1000 - paramMode1*100 - opCode) / 10000
	}

	return opCode, paramMode1, paramMode2, paramMode3
}

func add(numbers []int, position int, val1Mode int, val2Mode int) []int {
	val1 := 0
	if val1Mode == 1 {
		val1 = numbers[position+1]
	} else {
		val1 = numbers[numbers[position+1]]
	}
	val2 := 0
	if val2Mode == 1 {
		val2 = numbers[position+2]
	} else {
		val2 = numbers[numbers[position+2]]
	}

	numbers[numbers[position+3]] = val1 + val2

	return numbers
}

func multiply(numbers []int, position int, val1Mode int, val2Mode int) []int {
	val1 := 0
	if val1Mode == 1 {
		val1 = numbers[position+1]
	} else {
		val1 = numbers[numbers[position+1]]
	}
	val2 := 0
	if val2Mode == 1 {
		val2 = numbers[position+2]
	} else {
		val2 = numbers[numbers[position+2]]
	}

	numbers[numbers[position+3]] = val1 * val2
	return numbers
}

func jumpIfTrue(numbers []int, position int, paramMode1 int, paramMode2 int) (bool, int) {
	val1 := 0
	if paramMode1 == 1 {
		val1 = numbers[position+1]
	} else {
		val1 = numbers[numbers[position+1]]
	}

	if val1 == 0 {
		return false, 0
	}
	val2 := 0
	if paramMode2 == 1 {
		val2 = numbers[position+2]
	} else {
		val2 = numbers[numbers[position+2]]
	}

	return true, val2
}

func jumpIfFalse(numbers []int, position int, paramMode1 int, paramMode2 int) (bool, int) {
	val1 := 0
	if paramMode1 == 1 {
		val1 = numbers[position+1]
	} else {
		val1 = numbers[numbers[position+1]]
	}

	if val1 != 0 {
		return false, 0
	}
	val2 := 0
	if paramMode2 == 1 {
		val2 = numbers[position+2]
	} else {
		val2 = numbers[numbers[position+2]]
	}

	return true, val2
}

func lessThan(numbers []int, position int, paramMode1 int, paramMode2 int, paramMode3 int) []int {
	val1 := 0
	if paramMode1 == 1 {
		val1 = numbers[position+1]
	} else {
		val1 = numbers[numbers[position+1]]
	}

	val2 := 0
	if paramMode2 == 1 {
		val2 = numbers[position+2]
	} else {
		val2 = numbers[numbers[position+2]]
	}

	newValue := 0
	if val1 < val2 {
		newValue = 1
	}
	numbers[numbers[position+3]] = newValue

	return numbers
}

func equals(numbers []int, position int, paramMode1 int, paramMode2 int, paramMode3 int) []int {
	val1 := 0
	if paramMode1 == 1 {
		val1 = numbers[position+1]
	} else {
		val1 = numbers[numbers[position+1]]
	}

	val2 := 0
	if paramMode2 == 1 {
		val2 = numbers[position+2]
	} else {
		val2 = numbers[numbers[position+2]]
	}

	newValue := 0
	if val1 == val2 {
		newValue = 1
	}

	numbers[numbers[position+3]] = newValue

	return numbers
}

type Amp struct {
	CurrentPos int
	Numbers    []int
	Done       bool
	LastOutput int
	Phase      int
	Running    bool
}

func (amp *Amp) run(input int) {
	for {
		instruction := amp.Numbers[amp.CurrentPos]
		opCode, paramMode1, paramMode2, paramMode3 := interpretInstruction(instruction)

		if opCode == 99 {
			amp.Done = true
			return
		} else if opCode == 1 {
			amp.Numbers = add(amp.Numbers, amp.CurrentPos, paramMode1, paramMode2)
			amp.CurrentPos += 4
		} else if opCode == 2 {
			amp.Numbers = multiply(amp.Numbers, amp.CurrentPos, paramMode1, paramMode2)
			amp.CurrentPos += 4
		} else if opCode == 3 {
			if amp.Running == false {
				amp.Numbers[amp.Numbers[amp.CurrentPos+1]] = amp.Phase
				amp.Running = true
			} else {
				amp.Numbers[amp.Numbers[amp.CurrentPos+1]] = input
			}
			amp.CurrentPos += 2
		} else if opCode == 4 {
			if paramMode1 == 1 {
				amp.LastOutput = amp.Numbers[amp.CurrentPos+1]
			} else {
				amp.LastOutput = amp.Numbers[amp.Numbers[amp.CurrentPos+1]]
			}
			amp.CurrentPos += 2

			return
		} else if opCode == 5 {
			jump, newPosition := jumpIfTrue(amp.Numbers, amp.CurrentPos, paramMode1, paramMode2)
			if jump {
				amp.CurrentPos = newPosition
			} else {
				amp.CurrentPos += 3
			}
		} else if opCode == 6 {
			jump, newPosition := jumpIfFalse(amp.Numbers, amp.CurrentPos, paramMode1, paramMode2)
			if jump {
				amp.CurrentPos = newPosition
			} else {
				amp.CurrentPos += 3
			}
		} else if opCode == 7 {
			amp.Numbers = lessThan(amp.Numbers, amp.CurrentPos, paramMode1, paramMode2, paramMode3)
			amp.CurrentPos += 4
		} else if opCode == 8 {
			amp.Numbers = equals(amp.Numbers, amp.CurrentPos, paramMode1, paramMode2, paramMode3)
			amp.CurrentPos += 4
		}

	}
}

func initRun() int {

	max := 0

	for ampACounter := 5; ampACounter <= 9; ampACounter++ {

		for ampBCounter := 5; ampBCounter <= 9; ampBCounter++ {

			if ampBCounter == ampACounter {
				continue
			}

			for ampCCounter := 5; ampCCounter <= 9; ampCCounter++ {

				if ampCCounter == ampACounter || ampCCounter == ampBCounter {
					continue
				}

				for ampDCounter := 5; ampDCounter <= 9; ampDCounter++ {

					if ampDCounter == ampACounter || ampDCounter == ampBCounter || ampDCounter == ampCCounter {
						continue
					}

					for ampECounter := 5; ampECounter <= 9; ampECounter++ {
						if ampECounter == ampACounter || ampECounter == ampBCounter || ampECounter == ampCCounter || ampECounter == ampDCounter {
							continue
						}

						ampA := &Amp{
							CurrentPos: 0,
							Numbers:    getData(),
							Done:       false,
							Phase:      ampACounter,
						}

						ampB := &Amp{
							CurrentPos: 0,
							Numbers:    getData(),
							Done:       false,
							Phase:      ampBCounter,
						}

						ampC := &Amp{
							CurrentPos: 0,
							Numbers:    getData(),
							Done:       false,
							Phase:      ampCCounter,
						}

						ampD := &Amp{
							CurrentPos: 0,
							Numbers:    getData(),
							Done:       false,
							Phase:      ampDCounter,
						}

						ampE := &Amp{
							CurrentPos: 0,
							Numbers:    getData(),
							Done:       false,
							Phase:      ampECounter,
						}

						for {

							ampA.run(ampE.LastOutput)
							ampB.run(ampA.LastOutput)
							ampC.run(ampB.LastOutput)
							ampD.run(ampC.LastOutput)
							ampE.run(ampD.LastOutput)

							if ampA.Done || ampB.Done || ampC.Done || ampD.Done || ampE.Done {
								if ampE.LastOutput > max {
									max = ampE.LastOutput
								}
								break
							}
						}
					}
				}
			}
		}

	}

	return max
}

func main() {

	max := initRun()

	fmt.Printf("Signal to Thrusters %d\n", max)

}
