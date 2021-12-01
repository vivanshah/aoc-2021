package day

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Day18 struct {
	Equations []string
	Matches   map[int][]int
}

// ReadFile reads a file
func (d *Day18) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day18.txt")
	if err != nil {
		panic(err)
	}
	d.Equations = []string{}
	lines := strings.Split(string(input), "\r\n")
	for _, l := range lines {
		l = strings.TrimSpace(l)
		d.Equations = append(d.Equations, l)
	}

	fmt.Println("Read ", len(lines), " Lines")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day18) Part1() {
	fmt.Println("Day 18 Part 1")

	o := 0
	for _, line := range d.Equations {
		o += evaluate(line)
	}
	fmt.Println(o)
}

func evaluate(eq string) int {
	idxs := findBracketIdxs(eq)

	if len(idxs) == 0 {
		o := 0

		op := ""
		for _, s := range strings.Split(eq, " ") {
			if s == "+" || s == "*" {
				op = s
			} else {
				i, _ := strconv.Atoi(s)
				if op == "+" {
					o += i
				} else if op == "*" {
					o *= i
				} else {
					o = i
				}
				op = ""
			}
		}

		return o
	}

	subExpression := eq[idxs[0]+1 : idxs[1]]
	nextExpression := eq[:idxs[0]] + strconv.Itoa(evaluate(subExpression)) + eq[idxs[1]+1:]
	return evaluate(nextExpression)
}

func findBracketIdxs(expression string) []int {
	o := []int{}
	c := 0
	for idx, char := range expression {
		if char == '(' {
			c++
			if len(o) == 0 {
				o = append(o, idx)
			}
		} else if char == ')' {
			c--
			if c == 0 {
				o = append(o, idx)
				break
			}
		}
	}
	return o
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day18) Part2() {
	fmt.Println("Day 18 Part 2")
	o := 0
	for _, line := range d.Equations {
		o += evaluateAdvanced(line)
	}
	fmt.Println(o)
}

func evaluateAdvanced(expression string) int {
	idxs := findBracketIdxs(expression)

	if len(idxs) == 0 {
		idxs := findPlusIdxs(expression)

		if len(idxs) == 0 {
			o := 1
			for _, s := range strings.Split(expression, " ") {
				if s == "*" {
					continue
				}
				m, _ := strconv.Atoi(s)
				o *= m
			}
			return o
		}

		subExpression := expression[idxs[0] : idxs[1]+1]
		nextExpression := expression[:idxs[0]] + strconv.Itoa(evaluate(subExpression)) + expression[idxs[1]+1:]
		return evaluateAdvanced(nextExpression)
	}

	subExpression := expression[idxs[0]+1 : idxs[1]]
	nextExpression := expression[:idxs[0]] + strconv.Itoa(evaluateAdvanced(subExpression)) + expression[idxs[1]+1:]
	return evaluateAdvanced(nextExpression)
}

func findPlusIdxs(expression string) []int {
	plusIdx := -1
	for idx, char := range expression {
		if char == '+' {
			plusIdx = idx
			break
		}
	}
	if plusIdx < 0 {
		return []int{}
	}
	lower, upper := plusIdx-2, plusIdx+2
	for {
		if lower == 0 {
			break
		} else if expression[lower] == ' ' {
			lower++
			break
		} else {
			lower--
		}
	}
	for {
		if upper == len(expression)-1 {
			break
		} else if expression[upper] == ' ' {
			upper--
			break
		} else {
			upper++
		}
	}
	return []int{lower, upper}
}
