package day

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Day6 struct {
	Groups []Group
}

type Group struct {
	UniqueAnswers int
	Members       int
	AllAnswers    int
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day6) ReadFile(path string) error {
	fmt.Println("Reading input")
	file, err := os.Open("../../day6.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	d.Groups = []Group{}
	g := Group{}
	seen := map[rune]int{}
	for {
		line, err = reader.ReadString('\n')

		line = strings.TrimSpace(line)
		if len(line) > 0 {
			g.Members = g.Members + 1
		}
		//	fmt.Println(line)
		for _, c := range line {
			if seen[c] == 0 {
				g.UniqueAnswers++
			}
			seen[c] = seen[c] + 1
		}
		if err != nil && err == io.EOF {
			for _, v := range seen {
				if v == g.Members {
					g.AllAnswers = g.AllAnswers + 1
				}
			}
			//fmt.Println(seen)
			//	fmt.Println(g)
			d.Groups = append(d.Groups, g)
			break
		}
		if len(line) == 0 {
			for _, v := range seen {
				if v == g.Members {
					g.AllAnswers = g.AllAnswers + 1
				}
			}
			//	fmt.Println(seen)
			//	fmt.Println(g)
			d.Groups = append(d.Groups, g)
			g = Group{}
			seen = map[rune]int{}
			continue
		}

		if err != nil {
			for _, v := range seen {
				if v == g.Members {
					g.AllAnswers = g.AllAnswers + 1
				}
			}
			//	fmt.Println(seen)
			//	fmt.Println(g)
			d.Groups = append(d.Groups, g)
			//fmt.Println(err.Error())
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	fmt.Println("Read ", len(d.Groups), " Groups")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day6) Part1() {
	fmt.Println("Day 6 Part 1")
	sum := 0
	for _, g := range d.Groups {
		sum = sum + g.UniqueAnswers
	}
	fmt.Println(sum)

}

// Part2 executes part 2 of of this day's puzzle
func (d *Day6) Part2() {
	fmt.Println("Day 6 Part 2")
	sum := 0
	for _, g := range d.Groups {
		sum = sum + g.AllAnswers
	}
	fmt.Println(sum)
}
