package day

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Day13 struct {
	Earliest int
	Buses    []int
	AllBuses []string
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day13) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day13.txt")
	if err != nil {
		panic(err)
	}
	d.AllBuses = []string{}
	d.Buses = []int{}
	lines := strings.Split(string(input), "\r\n")
	//d.Rules = map[string][]string{}
	for i, l := range lines {
		l = strings.TrimSpace(l)
		if i == 0 {
			d.Earliest, _ = strconv.Atoi(l)
		} else {
			parts := strings.Split(l, ",")
			for _, p := range parts {
				d.AllBuses = append(d.AllBuses, p)
				if p != "x" {
					b, _ := strconv.Atoi(p)
					d.Buses = append(d.Buses, b)
				}
			}
		}

	}
	fmt.Println(d.Buses)
	fmt.Println("Read ", len(lines), " Instructions")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day13) Part1() {
	fmt.Println("Day 13 Part 1")
	m := map[int]int{}

	for _, b := range d.Buses {
		m[b] = b - (d.Earliest % b)
	}
	fmt.Println(m)
	for i := 939; i < math.MaxInt32; i++ {
		for _, b := range d.Buses {
			s := m[b]
			m[b] = m[b] - 1
			if s == 0 {
				fmt.Println("You can take bus: ", b, i)
				fmt.Println(b * (i - 939))
				return
			}
		}
	}

}

// Part2 executes part 2 of of this day's puzzle
func (d *Day13) Part2() {
	fmt.Println("Day 13 Part 2")

	e := map[int]int{}
	maxBus := 0
	fmt.Println(d.AllBuses)
	for i, b := range d.AllBuses {
		if b == "x" {
			continue
		}
		x, _ := strconv.Atoi(b)
		e[x] = i
		if x > maxBus {
			maxBus = x
		}
	}

	var found bool
	mult := 1
	f := map[int]bool{}
	for i := 0; i < math.MaxInt64; i = i + mult {
		found = true
		for k, v := range e {
			if (i+v)%k != 0 {
				found = false
				break
			}
			if !f[k] {
				mult = mult * k
			}
			f[k] = true
		}
		if found {
			fmt.Println(i)
			return
		}
	}
}
