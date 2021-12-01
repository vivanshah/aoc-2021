package day

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Day19 struct {
	Rules map[int]rule19
	Text  []string
}
type rule string

type message string
type rule19 struct {
	id   int
	text string
	rule []int
	alt  []int
}

var (
	regexRule = regexp.MustCompile("(\\d+): (.*?)$")
)

// ReadFile reads a file
func (d *Day19) ReadFile(path string) error {
	fmt.Println("Reading input")
	/*	input, err := ioutil.ReadFile("../../day19.txt")
		if err != nil {
			panic(err)
		}
		d.Rules = map[int]rule19{}
		sections := strings.Split(string(input), "\r\n\r\n")
		rules := strings.Split(sections[0], "\r\n")
		text := strings.Split(sections[1], "\r\n")
		for _, r := range rules {

			r = strings.ReplaceAll(strings.TrimSpace(r), ":", "")
			p := strings.Fields(r)
			fmt.Println(p)
			id, _ := strconv.Atoi(p[0])
			rule := rule19{id: id}
			if len(p) == 2 {
				c := p[1][1]
				rule.text = string(c)
				continue
			}

			rule.rule = []int{GetInt(p[1]), GetInt(p[2])}
			if len(p) > 3 {
				x := 4
				rule.alt = []int{}
				for x < len(p) {
					rule.alt = append(rule.alt)
					x++
				}

			}
			d.Rules[id] = rule
		}

		d.Text = []string{}
		for _, t := range text {
			d.Text = append(d.Text, t)
		}

		return nil*/
	input, err := ioutil.ReadFile("../../day19.txt")
	if err != nil {
		panic(err)
	}
	d.Rules = map[int]rule19{}
	inputParts := strings.Split(string(input), "\r\n\r\n")

	rules := make(map[int]rule)
	for _, ruleInput := range strings.Split(inputParts[0], "\r\n") {
		result := regexRule.FindStringSubmatch(ruleInput)
		newRuleIndex, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}
		newRule := rule(result[2])
		rules[newRuleIndex] = newRule
	}

	var messages []message
	for _, m := range strings.Split(inputParts[1], "\r\n") {
		messages = append(messages, message(m))
	}

	part1 := d.countRule0(rules, messages)
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"
	part2 := d.countRule0(rules, messages)

	fmt.Printf("1: %v\n2: %v\n", part1, part2)
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day19) Part1() {
	fmt.Println("Day 19 Part 1")

}

func (d *Day19) EvaluateRule(s string, r rule19) (string, bool) {
	if r.text == s {
		return "", true
	}
	s1, p1 := d.EvaluateRule(s, d.Rules[r.rule[0]])
	if p1 {
		s2, p2 := d.EvaluateRule(s1, d.Rules[r.rule[1]])
		return s2, p2
	} else {
		//alt
		s1, p1 := d.EvaluateRule(s, d.Rules[r.alt[0]])
		if p1 {
			s2, p2 := d.EvaluateRule(s1, d.Rules[r.alt[1]])
			return s2, p2
		}
	}
	return s, false
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day19) Part2() {
	fmt.Println("Day 19 Part 2")

}

func GetInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func (d Day19) countRule0(rules map[int]rule, messages []message) int {
	var checkMessageForRule func(message, int) (bool, []int)
	checkMessageForRule = func(m message, ruleNumber int) (bool, []int) {
		rule := rules[ruleNumber]
		switch rule[0] {
		case '"':
			if len(m) == 0 {
				return false, nil
			}
			if m[0] == rule[1] {
				return true, []int{1}
			}
			return false, nil
		default:
			orGroups := strings.Split(string(rule), " | ")
			var totalMatchesByAllOrGroups []int
			for _, orGroup := range orGroups {
				rulesInOrgroup := strings.Split(orGroup, " ")
				allCurrentIndex := []int{0}
				for _, fr := range rulesInOrgroup {
					nextRule, err := strconv.Atoi(fr)
					if err != nil {
						panic(err)
					}
					oldIndexLength := len(allCurrentIndex)
					for possibilityIndex := 0; possibilityIndex < oldIndexLength; possibilityIndex++ {
						currentIndex := allCurrentIndex[possibilityIndex]
						ruleApplies, length := checkMessageForRule(m[currentIndex:], nextRule)
						if ruleApplies {
							for _, newIndex := range length {
								if newIndex <= len(m) {
									allCurrentIndex = append(allCurrentIndex, currentIndex+newIndex)
								}
							}
						}
						allCurrentIndex = append(allCurrentIndex[:possibilityIndex], allCurrentIndex[possibilityIndex+1:]...)
						possibilityIndex--
						oldIndexLength--
					}
				}
				totalMatchesByAllOrGroups = append(totalMatchesByAllOrGroups, allCurrentIndex...)
			}
			if len(totalMatchesByAllOrGroups) > 0 {
				return true, totalMatchesByAllOrGroups
			} else {
				return false, nil
			}
		}
	}

	count := 0
	for _, m := range messages {
		if does, length := checkMessageForRule(m, 0); does {
			for _, l := range length {
				if l == len(m) {
					count++
					break
				}
			}

		}
	}
	return count
}
