package day

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Day5 struct {
	Passes []Pass
}

type Pass struct {
	Seat string
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day5) ReadFile(path string) error {
	fmt.Println("Reading input")
	file, err := os.Open("../../day5.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	d.Passes = []Pass{}
	for {
		line, err = reader.ReadString('\n')

		if err != nil && err != io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		d.Passes = append(d.Passes, Pass{Seat: line})

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	fmt.Println("Read ", len(d.Passes), " boarding passes")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day5) Part1() {
	fmt.Println("Day 5 Part 1")
	max := 0
	for _, p := range d.Passes {
		row, col := GetSeat(p)
		id := GetSeatID(row, col)
		if id > max {
			max = id
		}
	}
	fmt.Println(max)

}

func GetSeat(p Pass) (int, int) {
	var row, col int
	minRow := 0
	maxRow := 127
	minCol := 0
	maxCol := 7
	for i, c := range p.Seat {

		if i < 7 {
			//row
			if string(c) == "F" {
				maxRow = minRow + ((maxRow - minRow) / 2)
				row = maxRow
			} else if string(c) == "B" {
				minRow = maxRow - ((maxRow - minRow) / 2)
				row = minRow
			}
		} else {
			//column
			if string(c) == "L" {
				maxCol = minCol + ((maxCol - minCol) / 2)
				col = maxCol
			} else if string(c) == "R" {
				minCol = maxCol - ((maxCol - minCol) / 2)
				col = minCol
			}
		}
	}
	return row, col
}

// binary parsing version, quite a bit slower
/*func GetSeat(p Pass) (int, int) {
	r := strings.ReplaceAll(strings.ReplaceAll(p.Seat[:7], "F", "0"), "B", "1")
	row, _ := strconv.ParseInt(r, 2, 32)
	c := strings.ReplaceAll(strings.ReplaceAll(p.Seat[7:], "L", "0"), "R", "1")
	col, _ := strconv.ParseInt(c, 2, 32)
	return int(row), int(col)
}*/
func GetSeatID(row, col int) int {
	return (row * 8) + col
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day5) Part2() {
	fmt.Println("Day 5 Part 1")
	ids := map[int]bool{}
	for _, p := range d.Passes {
		row, col := GetSeat(p)
		if row == 0 || row == 127 {
			continue
		}
		id := GetSeatID(row, col)
		ids[id] = true
	}

	for k, v := range ids {
		if v && ids[k+2] && !ids[k+1] {
			fmt.Println(k + 1)
			break
		}

		if v && ids[k-2] && !ids[k-1] {
			fmt.Println(k - 1)
			break
		}
	}

}
