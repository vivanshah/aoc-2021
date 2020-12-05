package day

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Day2 struct {
	Entries []PasswordEntry
}

type PasswordEntry struct {
	Password string
	Min      int
	Max      int
	Letter   rune
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day2) ReadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	d.Entries = []PasswordEntry{}
	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		line = strings.TrimSuffix(line, "\n")
		entry := PasswordEntry{}
		fmt.Sscanf(line, "%d-%d %c: %s", &entry.Min, &entry.Max, &entry.Letter, &entry.Password)
		d.Entries = append(d.Entries, entry)
		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	fmt.Println("Finished reading input")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day2) Part1() {
	fmt.Println("Day 2 Part 1")
	valid := 0
	for _, e := range d.Entries {
		i := 0
		for _, c := range e.Password {
			if c == e.Letter {
				i++
			}
		}
		if i >= e.Min && i <= e.Max {
			valid++
		}

	}
	fmt.Println(valid)
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day2) Part2() {
	fmt.Println("Day 2 Part 2")
	valid := 0
	for _, e := range d.Entries {
		p1 := rune(e.Password[e.Min-1]) == e.Letter
		p2 := rune(e.Password[e.Max-1]) == e.Letter
		if (p1 || p2) && !(p1 && p2) {
			valid++
		}
	}
	fmt.Println(valid)
}
