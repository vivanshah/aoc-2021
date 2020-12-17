package day

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Day10 struct {
	Adapters []int
	Max      int
	aM       map[int][]int
	Target   int
	Ways     map[int]int
}

type Adapter struct {
	Output     int
	MinInput   int
	Downstream []Adapter
}

// ReadFile reads a file
func (d *Day10) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day10.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	d.Adapters = []int{}
	d.Adapters = append(d.Adapters, 0)
	max := 0
	for _, l := range lines {
		l = strings.TrimSpace(l)
		i, _ := strconv.Atoi(l)
		if i > max {
			max = i
		}
		d.Adapters = append(d.Adapters, i)
	}
	d.Target = max + 3
	d.Adapters = append(d.Adapters, max+3)
	sort.Ints(d.Adapters)
	fmt.Println("Read ", len(lines), " Joltage Adapters")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day10) Part1() {
	fmt.Println("Day 10 Part 1")
	chain := []int{}
	distribution := map[int]int{}
	joltage := 0
	for i := 0; i < len(d.Adapters); i++ {
		chain = append(chain, d.Adapters[i])
		distribution[d.Adapters[i]-joltage]++
		joltage = d.Adapters[i]
	}
	fmt.Println(distribution[1] * distribution[3])
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day10) Part2() {
	fmt.Println("Day 10 Part 2")
	//fmt.Println("Target: ", d.Target)
	d.aM = map[int][]int{}
	for _, i := range d.Adapters {
		if d.aM[i] == nil {
			d.aM[i] = []int{}
		}
		for j := 1; j <= 3; j++ {
			if d.aM[i-j] != nil {
				d.aM[i-j] = append(d.aM[i-j], i)
			}
		}

	}
	d.Ways = map[int]int{}
	fmt.Println(d.Count(d.Adapters[0], d.Target))
}

func (d *Day10) Count(output int, target int) int {
	c := 0
	if ways, exists := d.Ways[output]; exists {
		return ways
	}
	for _, o := range d.aM[output] {
		if o == target {
			c++
			continue
		}
		c += d.Count(o, target)
	}
	d.Ways[output] = c
	return c
}
