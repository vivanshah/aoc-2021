package day

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Day8 struct {
	Instructions []string
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day8) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day8.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	//d.Rules = map[string][]string{}
	d.Instructions = []string{}
	for _, l := range lines {
		l = strings.TrimSpace(l)
		d.Instructions = append(d.Instructions, l)

	}
	fmt.Println("Read ", len(lines), " Instructions")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day8) Part1() {
	fmt.Println("Day 8 Part 1")
	var acc int64
	acc = 0
	seen := map[int]bool{}
	seenSwappable := map[int]bool{}
	for x := 0; x < len(d.Instructions); x++ {
		if seen[x] {
			break
		}

		parts := strings.Fields(d.Instructions[x])
		instruction := parts[0]
		switch instruction {
		case "acc":
			seen[x] = true
			b, _ := strconv.Atoi(parts[1])
			acc = acc + int64(b)
		case "jmp":
			seen[x] = true
			seenSwappable[x] = true
			b, _ := strconv.Atoi(parts[1])
			x = x + b - 1
		case "nop":
			seen[x] = true
			seenSwappable[x] = true
			continue
		}
	}
	fmt.Println(acc)
	fmt.Println("Day 8 Part 2")
	for k := range seenSwappable {
		s2 := map[int]bool{}
		acc := 0
		loop := false
		for x := 0; x < len(d.Instructions); x++ {
			if s2[x] {
				loop = true
				break
			}
			s2[x] = true
			parts := strings.Fields(d.Instructions[x])
			instruction := parts[0]
			switch instruction {
			case "acc":
				b, _ := strconv.Atoi(parts[1])
				acc = acc + b
			case "jmp":
				if x == k {
					continue
				}
				b, _ := strconv.Atoi(parts[1])
				x = x + b - 1
			case "nop":
				if x != k {
					continue
				}
				b, _ := strconv.Atoi(parts[1])
				x = x + b - 1
			}

		}
		if !loop {
			fmt.Println(acc)
			break
		}
	}
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day8) Part2() {

	//see part 1
}
