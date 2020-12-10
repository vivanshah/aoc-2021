package day

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Day9 struct {
	Input []int
}

// ReadFile reads a file
func (d *Day9) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day9.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	d.Input = []int{}
	for _, l := range lines {
		l = strings.TrimSpace(l)
		i, _ := strconv.Atoi(l)
		d.Input = append(d.Input, i)
	}
	fmt.Println("Read ", len(lines), " Lines")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day9) Part1() {
	fmt.Println("Day 9 Part 1")
	preambleSize := 25
	for x := preambleSize; x < len(d.Input); x++ {
		if !d.VerifyValue(x, preambleSize) {
			fmt.Println(d.Input[x])
			break
		}
	}
}

func (d *Day9) VerifyValue(position int, preambleSize int) bool {
	toCheck := d.Input[position]
	seen := map[int]bool{}
	for x := position - preambleSize; x < position; x++ {
		t := toCheck - d.Input[x]
		_, exists := seen[t]
		if exists {
			return true
		}
		seen[d.Input[x]] = true
	}
	return false
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day9) Part2() {
	fmt.Println("Day 9 Part 2")

	target := 257342611
	i := 0
	j := 1
	min := math.MaxInt32
	max := 0
	acc := 0
	acc = d.Input[i] + d.Input[j]
	for acc != target && i < len(d.Input) && j < len(d.Input)-1 {
		for acc < target {
			j++
			acc += d.Input[j]
		}
		for acc > target {
			acc -= d.Input[i]
			i++
		}
	}
	for x := i; x <= j; x++ {
		if d.Input[x] < min {
			min = d.Input[x]
		} else if d.Input[x] > max {
			max = d.Input[x]
		}
	}

	fmt.Println(min + max)
}
