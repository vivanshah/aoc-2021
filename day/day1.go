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

// Part1 executs part 1 of of this day's puzzle
func (d *Day1) Part1() {
	fmt.Println("Day 1 Part 1")
	err := d.ReadFile("1-1.txt")
	if err != nil {
		panic(err)
	}

	m := map[int]bool{}
	for _, e := range d.Entries {
		m[e] = true
		if m[2020-e] == true {
			fmt.Printf("%d * %d = %d\n", e, 2020-e, e*(2020-e))
		}
	}
}

// Part2 executs part 2 of of this day's puzzle
func (d *Day1) Part2() {
	fmt.Println("Day 1 Part 2")
	err := d.ReadFile("1-1.txt")
	if err != nil {
		panic(err)
	}
	for _, i := range d.Entries {
		t := 2020 - i
		m := map[int]bool{}
		for _, e := range d.Entries {
			m[e] = true
			if m[t-e] == true {
				fmt.Printf("%d * %d * %d = %d\n", e, t-i, i, i*e*(t-e))
				return
			}
		}
	}

}
