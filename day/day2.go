package day

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Day2 struct {
	Commands []Command
}

type Command struct {
	Direction string
	Distance  int
}

func (d *Day2) GetDayNumber() int {
	return 2
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day2) ReadFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	d.Commands = []Command{}
	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		var c Command
		_, err := fmt.Sscanf(line, "%s %d", &c.Direction, &c.Distance)
		d.Commands = append(d.Commands, c)
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
func (d *Day2) Part1() {
	fmt.Println("Day 2 Part 1")
	var position, depth int
	for _, c := range d.Commands {
		fmt.Println(c.Direction, c.Distance)
		switch c.Direction {
		case "forward":
			position += c.Distance
		case "down":
			depth += c.Distance
		case "up":
			depth -= c.Distance

		}
	}
	fmt.Println(position * depth)
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day2) Part2() {
	fmt.Println("Day 2 Part 2")
	var position, depth, aim int
	for _, c := range d.Commands {
		fmt.Println(c.Direction, c.Distance)
		switch c.Direction {
		case "forward":
			position += c.Distance
			depth = depth + aim*c.Distance
		case "down":
			aim += c.Distance
		case "up":
			aim -= c.Distance
		}
	}
	fmt.Printf("Position: %d Depth: %d\n", position, depth)
	fmt.Println(position * depth)
}
