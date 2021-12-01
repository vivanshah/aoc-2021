package day

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Day17 struct {
	Cubes          map[point]int
	Visiblepoints  map[point]map[point]bool
	Adjacentpoints map[point]map[point]bool
	Lines          []string
	ActiveCubes    map[point]bool
}

type point struct {
	x int
	y int
	z int
	w int
}

// ReadFile reads a file
func (d *Day17) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day17.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	d.Cubes = map[point]int{}
	d.Lines = lines
	d.ActiveCubes = map[point]bool{}
	for x, l := range lines {
		l = strings.TrimSpace(l)
		for y, p := range l {
			c := point{x: x, y: y, z: 0, w: 0}
			switch string(p) {
			case ".":
				d.Cubes[c] = 0
			case "#":
				d.Cubes[c] = 1
				d.ActiveCubes[c] = true
			}
		}
	}

	fmt.Println("Read ", len(lines), " Lines")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day17) Part1() {
	fmt.Println("Day 17 Part 1")

	c := 0
	d.Adjacentpoints = map[point]map[point]bool{}
	mx := len(d.Lines)
	my := len(d.Lines[0])
	for x := -1; x <= mx; x++ {
		for y := -1; y <= my; y++ {
			for z := -1; z <= 1; z++ {
				p := point{x: x, y: y, z: z}
				_, ok := d.Cubes[p]
				if !ok {
					d.Cubes[p] = 0
				}
			}
		}

	}
	for c < 6 {
		pointsToActivate := []point{}
		pointsToDeactivate := []point{}

		for p := range d.Cubes {
			switch d.GetpointEvent(p) {
			case 1:
				pointsToActivate = append(pointsToActivate, p)
			case 0:
				pointsToDeactivate = append(pointsToDeactivate, p)
			}

		}

		for _, s := range pointsToDeactivate {
			d.Cubes[s] = 0
		}
		for _, s := range pointsToActivate {
			d.Cubes[s] = 1
		}

		c++
	}

	fmt.Println(d.CountActive())
}
func (d *Day17) CountActive() int {
	o := 0
	for _, v := range d.Cubes {
		if v == 1 {
			o++
		}
	}
	return o
}

func (d *Day17) GetpointEvent(s point) int {
	adjacentActive := 0
	adjacent, ok := d.Adjacentpoints[s]
	if ok {
		for k := range adjacent {
			if d.Cubes[k] == 1 {
				adjacentActive++
			}
		}
	} else {
		d.Adjacentpoints[s] = map[point]bool{}
		for x := s.x - 1; x <= s.x+1; x++ {
			for y := s.y - 1; y <= s.y+1; y++ {
				for z := s.z - 1; z <= s.z+1; z++ {
					if s.x == x && s.y == y && s.z == z {
						continue
					}
					c := point{x: x, y: y, z: z}
					_, ok := d.Cubes[c]
					if !ok {
						d.Cubes[c] = 0
					}

					if d.Cubes[c] == 1 { //active
						adjacentActive++
					}
					d.Adjacentpoints[s][c] = true
				}
			}
		}
	}
	if d.Cubes[s] == 1 && adjacentActive == 2 || adjacentActive == 3 {
		return 1
	}
	if d.Cubes[s] == 0 && adjacentActive == 3 {
		return 1
	}
	return 0
}

func (d *Day17) GetpointEvent2(s point) int {
	adjacentActive := 0
	adjacent, ok := d.Adjacentpoints[s]
	var c point
	if ok {
		for k := range adjacent {
			if d.Cubes[k] == 1 {
				adjacentActive++
			}
		}
	} else {
		d.Adjacentpoints[s] = map[point]bool{}
		for x := s.x - 1; x <= s.x+1; x++ {
			for y := s.y - 1; y <= s.y+1; y++ {
				for z := s.z - 1; z <= s.z+1; z++ {
					for w := s.w - 1; w <= s.w+1; w++ {
						if s.x == x && s.y == y && s.z == z && s.w == w {
							continue
						}
						c = point{x: x, y: y, z: z, w: w}
						_, ok := d.Cubes[c]
						if !ok {
							d.Cubes[c] = 0
						}

						if d.Cubes[c] == 1 { //active
							adjacentActive++
						}
						d.Adjacentpoints[s][c] = true
					}
				}
			}
		}
	}
	if d.Cubes[s] == 1 && adjacentActive == 2 || adjacentActive == 3 {
		return 1
	}
	if d.Cubes[s] == 0 && adjacentActive == 3 {
		return 1
	}
	return 0
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day17) Part2() {
	fmt.Println("Day 17 Part 2")
	c := 0
	d.Adjacentpoints = map[point]map[point]bool{}
	mx := len(d.Lines)
	my := len(d.Lines[0])
	for x := -1; x <= mx; x++ {
		for y := -1; y <= my; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					p := point{x: x, y: y, z: z, w: w}
					_, ok := d.Cubes[p]
					if !ok {
						d.Cubes[p] = 0
					}
				}
			}
		}

	}
	var cubeEvent int
	for c < 6 {
		pointsToActivate := []point{}
		pointsToDeactivate := []point{}

		for p := range d.Cubes {
			adjacentActive := 0
			cubeEvent = 0
			adjacent, ok := d.Adjacentpoints[p]
			var c point
			if ok {
				for k := range adjacent {
					if d.Cubes[k] == 1 {
						adjacentActive++
					}
				}
			} else {
				d.Adjacentpoints[p] = map[point]bool{}
				for x := p.x - 1; x <= p.x+1; x++ {
					for y := p.y - 1; y <= p.y+1; y++ {
						for z := p.z - 1; z <= p.z+1; z++ {
							for w := p.w - 1; w <= p.w+1; w++ {
								if p.x == x && p.y == y && p.z == z && p.w == w {
									continue
								}
								c = point{x: x, y: y, z: z, w: w}
								_, ok := d.Cubes[c]
								if !ok {
									d.Cubes[c] = 0
								}

								if d.Cubes[c] == 1 { //active
									adjacentActive++
								}
								d.Adjacentpoints[p][c] = true
							}
						}
					}
				}
			}
			if d.Cubes[p] == 1 && adjacentActive == 2 || adjacentActive == 3 {
				cubeEvent = 1
			}
			if d.Cubes[p] == 0 && adjacentActive == 3 {
				cubeEvent = 1
			}

			switch cubeEvent {
			case 1:
				pointsToActivate = append(pointsToActivate, p)
			case 0:
				pointsToDeactivate = append(pointsToDeactivate, p)
			}

		}

		for _, s := range pointsToDeactivate {
			d.Cubes[s] = 0
		}
		for _, s := range pointsToActivate {
			d.Cubes[s] = 1
		}

		c++
	}

	fmt.Println(d.CountActive())
}
