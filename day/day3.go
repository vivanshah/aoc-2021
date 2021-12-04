package day

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type Day3 struct {
	Diagnostic []int
	Width      int
}

func (d *Day3) GetDayNumber() int {
	return 3
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day3) ReadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	d.Diagnostic = []int{}
	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		line = strings.TrimSuffix(line, "\n")
		if d.Width == 0 {
			d.Width = len(line)
		}
		var r int
		_, err := fmt.Sscanf(line, "%b", &r)
		if err != nil {
			break
		}
		d.Diagnostic = append(d.Diagnostic, r)
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	fmt.Println("Width: ", d.Width)
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day3) Part1() {
	fmt.Println("Day 3 Part 1")
	l := len(d.Diagnostic)
	t := l / 2
	fmt.Println("Target: ", t)
	g := 0
	e := 0
	m := make([]int, d.Width)

	for i := 0; i < len(d.Diagnostic); i++ {
		r := d.Diagnostic[i]
		for c := 0; c < d.Width; c++ {
			if (r>>c)&1 == 1 {
				m[c]++
			}
		}
	}
	for i := len(m) - 1; i >= 0; i-- {
		s := m[i]
		if s > t {
			g = g | 1
		}
		if s < t {
			e = e | 1
		}
		g = g << 1
		e = e << 1
	}
	g = g >> 1
	e = e >> 1
	fmt.Printf("G: %d\n", g)
	fmt.Printf("E: %d\n", e)
	fmt.Println(e * g)
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day3) Part2() {
	fmt.Println("Day 3 Part 2")
	l := len(d.Diagnostic)

	o2discard := make([]int, len(d.Diagnostic))
	co2discard := make([]int, len(d.Diagnostic))
	for i := d.Width - 1; i >= 0; i-- {
		m := make([]int, d.Width)
		t := int(math.Ceil(float64(l) / float64(2)))
		for i := 0; i < len(d.Diagnostic); i++ {
			if o2discard[i] == 1 {
				continue
			}
			r := d.Diagnostic[i]
			for c := 0; c < d.Width; c++ {
				if (r>>c)&1 == 1 {
					m[c]++
				}
			}
		}
		s := m[i]

		for iD, r := range d.Diagnostic {
			if o2discard[iD] == 1 {
				continue
			}
			b := r >> i
			if s >= t {
				//more one bits
				if b&1 == 0 {
					l--
					o2discard[iD] = 1
				}
			} else if s < t {
				//more zero bits
				if b&1 == 1 {
					l--
					o2discard[iD] = 1
				}
			}
		}
		if l == 1 {
			break
		}
	}

	l = len(d.Diagnostic)
	for i := d.Width - 1; i >= 0; i-- {
		m := make([]int, d.Width)
		t := int(math.Ceil(float64(l) / float64(2)))
		for i := 0; i < len(d.Diagnostic); i++ {
			if co2discard[i] == 1 {
				continue
			}
			r := d.Diagnostic[i]
			for c := 0; c < d.Width; c++ {
				if (r>>c)&1 == 1 {
					m[c]++
				}
			}
		}
		s := m[i]

		for iD, r := range d.Diagnostic {
			if co2discard[iD] == 1 {
				continue
			}
			b := r >> i
			if s >= t {
				//fewer zero bits
				if b&1 == 1 {
					l--
					co2discard[iD] = 1
				}
			} else if s < t {
				//fewer one bits
				if b&1 == 0 {
					l--
					co2discard[iD] = 1
				}
			}

		}
		if l == 1 {
			break
		}
	}
	oRating := 0
	for i, o := range o2discard {
		if o == 0 {
			oRating = d.Diagnostic[i]
		}
	}
	fmt.Println("Oxygen Scrubber Rating:", oRating)
	co2Rating := 0
	for i, o := range co2discard {
		if o == 0 {
			co2Rating = d.Diagnostic[i]
		}
	}
	fmt.Println("C02 Scrubber Rating:", co2Rating)

	fmt.Println(oRating * co2Rating)
}
