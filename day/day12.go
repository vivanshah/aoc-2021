package day

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Day12 struct {
	Instructions []string
	Grid         map[coordinate]int
}

type coordinate struct {
	I int
	J int
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day12) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day12.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	//d.Rules = map[string][]string{}
	d.Instructions = []string{}
	d.Grid = map[coordinate]int{}
	for _, l := range lines {
		l = strings.TrimSpace(l)
		d.Instructions = append(d.Instructions, l)

	}
	fmt.Println("Read ", len(lines), " Instructions")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day12) Part1() {
	fmt.Println("Day 12 Part 1")
	var magnitude int
	orientation := "E"
	var location coordinate
	for x := 0; x < len(d.Instructions); x++ {
		instruction := d.Instructions[x]
		fmt.Println(instruction)
		direction := string(instruction[0])
		if len(instruction) > 1 {
			magnitude, _ = strconv.Atoi(instruction[1:])
		}
		switch direction {
		case "N":
			location.I -= magnitude
		case "S":
			location.I += magnitude
		case "E":
			location.J += magnitude
		case "W":
			location.J -= magnitude
		case "L":
			fmt.Printf("turning left %d from %s\n", magnitude, orientation)

			for magnitude > 0 {
				switch orientation {
				case "N":
					orientation = "W"
				case "S":
					orientation = "E"
				case "E":
					orientation = "N"
				case "W":
					orientation = "S"
				}
				magnitude -= 90
			}
			fmt.Printf("now facing %s\n", orientation)
		case "R":
			for magnitude > 0 {
				switch orientation {
				case "N":
					orientation = "E"
				case "S":
					orientation = "W"
				case "E":
					orientation = "S"
				case "W":
					orientation = "N"
				}
				magnitude -= 90
			}
		case "F":
			switch orientation {
			case "N":
				location.I -= magnitude
			case "S":
				location.I += magnitude
			case "E":
				location.J += magnitude
			case "W":
				location.J -= magnitude
			}
		}
	}
	fmt.Println(location.I, location.J)

}

// Part2 executes part 2 of of this day's puzzle
func (d *Day12) Part2() {
	fmt.Println("Day 12 Part 2")
	var magnitude int
	var location coordinate
	waypoint := coordinate{I: -1, J: 10}
	for x := 0; x < len(d.Instructions); x++ {
		instruction := d.Instructions[x]
		fmt.Println(instruction)
		direction := string(instruction[0])
		if len(instruction) > 1 {
			magnitude, _ = strconv.Atoi(instruction[1:])
		}
		fmt.Printf("magnitude %d. waypoint at %v\n", magnitude, waypoint)
		switch direction {
		case "N":
			waypoint.I -= magnitude
		case "S":
			waypoint.I += magnitude
		case "E":
			waypoint.J += magnitude
		case "W":
			waypoint.J -= magnitude
		case "L":
			for magnitude > 0 {
				i := waypoint.I
				waypoint.I = -1 * waypoint.J
				waypoint.J = i
				magnitude -= 90
			}
		case "R":
			for magnitude > 0 {
				j := waypoint.J
				waypoint.J = -1 * waypoint.I
				waypoint.I = j
				magnitude -= 90
			}
		case "F":
			location.I += magnitude * waypoint.I
			location.J += magnitude * waypoint.J
		}
		fmt.Println("ship at ", location, ", waypoint at ", waypoint)
	}
	fmt.Println(location.I, location.J)

}

func RotateWaypoint(w coordinate) coordinate {
	if w.I < 0 && w.J > 0 {
		i := w.I
		w.I = -1 * w.J
		w.J = i
	} else if w.I < 0 && w.J < 0 {
		i := w.I
		w.I = -1 * w.J
		w.J = i
	} else if w.I > 0 && w.J < 0 {
		i := w.I
		w.I = -1 * w.J
		w.J = i
	}
	return coordinate{}
}
