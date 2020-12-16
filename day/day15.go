package day

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Day15 struct {
	Starting []int
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day15) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day15.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	d.Starting = []int{}
	for _, l := range lines {
		l = strings.TrimSpace(l)
		parts := strings.Split(l, ",")
		for _, n := range parts {
			i, _ := strconv.Atoi(n)
			d.Starting = append(d.Starting, i)
		}
	}
	fmt.Println("Read ", len(d.Starting), " Starting Numbers")
	return nil
}

func (d *Day15) Part1() {
	fmt.Println("Day 15 Part 1")
	var lastSpoken int
	sp := make(map[int][]int, 30000000)
	for i, s := range d.Starting {
		lastSpoken = s
		if sp[s] == nil {
			sp[s] = []int{}
		}
		sp[s] = append(sp[s], i+1)
	}
	var s int
	var l []int
	var length int
	for turn := len(d.Starting) + 1; turn <= 2020; turn++ {
		l = sp[lastSpoken]
		length = len(l)
		if length == 1 {
			//previous utterance was the first time it was spoken
			s = 0
		} else if length > 1 {
			//spoken before this last time
			s = (turn - 1) - sp[lastSpoken][length-2]
		}
		lastSpoken = s
		if sp[s] == nil {
			sp[s] = []int{}
		}
		sp[s] = append(sp[s], turn)
	}

	fmt.Println(lastSpoken)
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day15) Part2() {
	fmt.Println("Day 15 Part 2")
	//var lastSpoken int
	next := 0
	turn := 0
	sp := make(map[int]int, 30000000)
	for ; turn < len(d.Starting); turn++ {
		next = d.Starting[turn]
		sp[next] = turn + 1
	}
	var m int
	var ok bool
	for ; turn < 30000000; turn++ {
		m, ok = sp[next]
		sp[next] = turn
		if !ok {
			next = 0
		} else {
			next = turn - m
		}
	}

	fmt.Println(next)
}
