package day

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Day16 struct {
	Fields        []field
	ValidValues   []int
	MyTicket      ticket
	NearbyTickets []ticket
}

type ticket struct {
	fieldValues []int
}

type field struct {
	Name   string
	Range1 fieldRange
	Range2 fieldRange
}
type fieldRange struct {
	Min int
	Max int
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day16) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day16.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	section := 1
	d.Fields = []field{}
	d.ValidValues = make([]int, 10000000)
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			section++
			continue
		}
		if section == 1 {
			parts := strings.Split(l, ": ")
			ranges := strings.Split(parts[1], " or ")
			range1 := strings.Split(ranges[0], "-")
			range2 := strings.Split(ranges[1], "-")
			r1min, _ := strconv.Atoi(range1[0])
			r1max, _ := strconv.Atoi(range1[1])
			r2min, _ := strconv.Atoi(range2[0])
			r2max, _ := strconv.Atoi(range2[1])
			f := field{Name: parts[0], Range1: fieldRange{Min: r1min, Max: r1max}, Range2: fieldRange{Min: r2min, Max: r2max}}
			d.Fields = append(d.Fields, f)

			for i := r1min; i <= r1max; i++ {
				d.ValidValues[i] = 1
			}
			for i := r2min; i <= r2max; i++ {
				d.ValidValues[i] = 1
			}
		}
		if section == 2 {
			if l == "your ticket:" {
				continue
			}
			t := ticket{fieldValues: []int{}}
			parts := strings.Split(l, ",")
			for _, p := range parts {
				v, _ := strconv.Atoi(p)
				t.fieldValues = append(t.fieldValues, v)
			}
			d.MyTicket = t
		}
		if section == 3 {
			if l == "nearby tickets:" {
				continue
			}
			t := ticket{fieldValues: []int{}}
			parts := strings.Split(l, ",")
			for _, p := range parts {
				v, _ := strconv.Atoi(p)
				t.fieldValues = append(t.fieldValues, v)
			}
			d.NearbyTickets = append(d.NearbyTickets, t)
		}

	}
	return nil
}

func (d *Day16) Part1() {
	fmt.Println("Day 16 Part 1")

	s := 0
	for _, t := range d.NearbyTickets {
		for _, v := range t.fieldValues {
			if d.ValidValues[v] == 0 {
				s += v
			}
		}
	}
	fmt.Println(s)

}

// Part2 executes part 2 of of this day's puzzle
func (d *Day16) Part2() {
	fmt.Println("Day 16 Part 2")
	validTickets := []ticket{}
	for _, t := range d.NearbyTickets {
		valid := true
		for _, v := range t.fieldValues {
			if d.ValidValues[v] == 0 {
				valid = false
				break
			}
		}
		if valid {
			validTickets = append(validTickets, t)
		}
	}
	c := len(validTickets[0].fieldValues)
	m := map[string]map[int]bool{}

	for i := 0; i < c; i++ {
		for _, f := range d.Fields {
			valid := true
			for _, t := range validTickets {
				v := t.fieldValues[i]
				if (f.Range1.Min <= v && f.Range1.Max >= v) || (f.Range2.Min <= v && f.Range2.Max >= v) {
					//fmt.Println(v, " is invalid for field ", f.Name, f.Range1, f.Range2)
					continue
				}
				valid = false
			}
			if valid { //all field values at index i were valid for this field, add it to candidate list
				//fmt.Println(f.Name, i)
				if m[f.Name] == nil {
					m[f.Name] = map[int]bool{}
				}
				m[f.Name][i] = true
			}
			//fmt.Println(m)
		}
	}
	result := 1
	//fmt.Println(d.MyTicket)
	for {
		i := 0
		for fieldName, candidates := range m {
			if len(candidates) == 0 {
				i++
				continue
			}
			if len(candidates) == 1 {
				for index := range candidates {
					if strings.HasPrefix(fieldName, "departure") {
						result = result * d.MyTicket.fieldValues[index]
					}
					for _, j := range m {
						_, ok := j[index]
						if ok {
							delete(j, index)
						}
					}
				}
			}

		}
		if i == len(m) {
			break
		}

	}
	fmt.Println(result)
}
