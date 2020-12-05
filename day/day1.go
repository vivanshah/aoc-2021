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
	m := map[int]bool{}
	for _, e := range d.Entries {
		m[e] = true
		if m[2020-e] {
			fmt.Printf("%d * %d = %d\n", e, 2020-e, e*(2020-e))
			return
		}
	}
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day1) Part2() {
	for _, i := range d.Entries {
		t := 2020 - i
		m := map[int]bool{}
		for _, e := range d.Entries {
			m[e] = true
			if m[t-e] {
				fmt.Printf("%d * %d * %d = %d\n", e, t-i, i, i*e*(t-e))
				return
			}
		}
	}
}
