package day

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Day7 struct {
	Rules      map[string]map[string]int
	Containers map[string][]string
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day7) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day7.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	//d.Rules = map[string][]string{}
	d.Containers = map[string][]string{}
	d.Rules = map[string]map[string]int{}
	for _, l := range lines {
		l = strings.TrimSpace(l)

		p := strings.Split(l, " bags contain ")
		var count int
		var color string
		for _, s := range strings.Split(p[1], " ") {
			if strings.HasSuffix(s, ",") || strings.HasSuffix(s, ".") {
				color = strings.TrimSpace(color)
				d.Containers[color] = append(d.Containers[color], p[0])
				if _, ok := d.Rules[p[0]]; !ok {
					d.Rules[p[0]] = map[string]int{}
				}
				d.Rules[p[0]][color] = count
				count = 0
				color = ""
				continue
			}
			i, err := strconv.Atoi(s)
			if err != nil {
				color = color + " " + s
				continue
			}
			count = i

		}

	}
	fmt.Println("Read ", len(lines), " Lines")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day7) Part1() {
	fmt.Println("Day 7 Part 1")
	stack := Stack{}
	stack.Push("shiny gold")
	bags := map[string]bool{}
	for {
		t, ok := stack.Pop()
		if !ok {
			break
		}
		v, ok := d.Containers[t]
		bags[t] = true
		for _, s := range v {
			stack.Push(s)
		}
	}
	fmt.Println(len(bags) - 1)

}

// Part2 executes part 2 of of this day's puzzle
func (d *Day7) Part2() {
	fmt.Println("Day 7 Part 2")
	stack := Stack{}
	stack.Push("shiny gold")
	seen := map[string]bool{}
	sum := 0
	for {
		t, ok := stack.Pop()
		if !ok {
			break
		}
		for k, v := range d.Rules[t] {
			sum = sum + v
			if _, ok := seen[k]; !ok {
				for x := 0; x < v; x++ {
					stack.Push(k)
				}
			}
		}
	}
	fmt.Println(sum)
}

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
