package day

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Day25 struct {
	cardPK int
	doorPK int
}

// ReadFile reads a file
func (d *Day25) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day25.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	for i, l := range lines {
		l = strings.TrimSpace(l)
		if i == 0 {
			d.cardPK = GetInt(l)
		}
		if i == 1 {
			d.doorPK = GetInt(l)
		}
	}
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day25) Part1() {
	fmt.Println("Day 25 Part 1")

	loop := 0
	for k := 1; k != d.cardPK; loop++ {
		k = k * 7 % 20201227
	}

	key := 1
	for l := 0; l < loop; l++ {
		key = key * d.doorPK % 20201227
	}
	fmt.Println(key)
	/*
			vC := 0
		vD := 0
		cardLoopSize := 1
		doorLoopSize := 1

			for {
				//	fmt.Println("Testing Card Loop Size: ", cardLoopSize)
				vC = transformSubject(7, cardLoopSize)
				if vC == d.cardPK {
					fmt.Println("Card Loop Size: ", cardLoopSize)
					break
				}
				cardLoopSize++
			}
			for {
				//	fmt.Println("Testing Door Loop Size: ", doorLoopSize)
				vD = transformSubject(7, doorLoopSize)
				if vD == d.doorPK {
					fmt.Println("Door Loop Size: ", doorLoopSize)
					break
				}
				doorLoopSize++
			}

			fmt.Println()
	*/
}

func transformSubject(subject int, loopSize int) int {
	v := 1
	for i := 0; i < loopSize; i++ {
		v = v * subject
		v = v % 20201227
	}
	return v
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day25) Part2() {
	fmt.Println("Day 25 Part 2")

	fmt.Println()
}
