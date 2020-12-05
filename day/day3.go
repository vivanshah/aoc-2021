package day

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Day3 struct {
	Trees [][]int
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day3) ReadFile(path string) error {
	file, err := os.Open("../../day3.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	var row, col int
	row = 0

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		line = strings.TrimSuffix(line, "\r\n")
		if len(line) == 0 {
			break
		}
		d.Trees = append(d.Trees, make([]int, len(line)))
		var c rune
		//	fmt.Sscanf(line, "%d-%d %c: %s", &entry.Min, &entry.Max, &entry.Letter, &entry.Password)
		for col, c = range line {
			if string(c) == "#" {
				d.Trees[row][col] = 1
			} else {
				d.Trees[row][col] = 0
			}
		}
		if err != nil {
			break
		}
		row++
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	fmt.Println("Finished reading input")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day3) Part1() {
	fmt.Println("Day 3 Part 1")
	fmt.Println(d.CountTreesForSlope(3, 1))
}

func (d *Day3) CountTreesForSlope(right, down int) int {
	var row, col int
	col = 0
	row = 0
	trees := 0
	fmt.Println(len(d.Trees), len(d.Trees[0]))
	for {

		if d.Trees[row][col] == 1 {
			trees++

		}
		fmt.Println("Row: ", row, " Col: ", col, " Tree: ", d.Trees[row][col])
		fmt.Println("Trees: ", trees)
		col += right
		row += down
		if col >= len(d.Trees[0]) {
			col = col % len(d.Trees[0])
		}
		if row >= len(d.Trees) {
			break
		}

	}
	return trees
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day3) Part2() {
	fmt.Println("Day 3 Part 2")
	fmt.Println(d.CountTreesForSlope(1, 1) * d.CountTreesForSlope(3, 1) * d.CountTreesForSlope(5, 1) * d.CountTreesForSlope(7, 1) * d.CountTreesForSlope(1, 2))
}
