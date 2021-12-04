package day

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Day1 struct {
	Entries []int
}

func (d *Day1) GetDayNumber() int {
	return 1
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day1) ReadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	d.Entries = []int{}
	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		i, _ := strconv.Atoi(line)
		d.Entries = append(d.Entries, i)
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day1) Part1() {
	fmt.Println("Day 1 Part 1")
	c := 0
	for i, e := range d.Entries {
		if i > 0 {
			if e > d.Entries[i-1] {
				c++
			}
		}
	}
	fmt.Println(c)
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day1) Part2() {
	fmt.Println("Day 1 Part 2")
	sums := []int{}
	for i := 2; i < len(d.Entries); i++ {

		s := d.Entries[i] + d.Entries[i-1] + d.Entries[i-2]
		sums = append(sums, s)

	}
	c := 0
	for i, v := range sums {
		fmt.Println(v)
		if i > 0 {
			if v > sums[i-1] {
				c++
			}
		}
	}
	fmt.Println(c)
}
