package day

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Day11 struct {
	Seats         map[Seat]int
	Height        int
	Width         int
	VisibleSeats  map[Seat]map[Seat]bool
	AdjacentSeats map[Seat]map[Seat]bool
}

type Seat struct {
	I int
	J int
}

// ReadFile reads a file
func (d *Day11) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day11.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	d.Seats = map[Seat]int{}
	d.Height = len(lines)
	for i, l := range lines {
		l = strings.TrimSpace(l)
		d.Width = len(l)
		for j, s := range l {
			switch s {
			case rune('L'):
				d.Seats[Seat{I: i, J: j}] = 1
			case rune('.'):
				d.Seats[Seat{I: i, J: j}] = 0
			}
		}
	}
	fmt.Println("Read ", len(lines), " Lines")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day11) Part1() {
	fmt.Println("Day 11 Part 1")
	c := 0
	d.AdjacentSeats = map[Seat]map[Seat]bool{}
	for {
		seatsToOccupy := []Seat{}
		seatsToVacate := []Seat{}
		for i := 0; i < d.Height; i++ {
			for j := 0; j < d.Width; j++ {
				s := Seat{I: i, J: j}
				if d.Seats[s] == 0 {
					continue
				}
				switch d.GetSeatEvent(s, 4) {
				case 1:
					seatsToVacate = append(seatsToVacate, s)
				case 2:
					seatsToOccupy = append(seatsToOccupy, s)
				}
			}
		}
		if len(seatsToOccupy) == 0 && len(seatsToVacate) == 0 {
			d.printGrid()
			d.CountOccupied()
			break
		}
		for _, s := range seatsToVacate {
			d.Seats[s] = 1
		}
		for _, s := range seatsToOccupy {
			d.Seats[s] = 2
		}
		if c%2 == 0 {
			d.printGrid()
			time.Sleep(100 * time.Millisecond)
		}
		c++
	}

}
func (d *Day11) CountOccupied() {
	o := 0
	for _, v := range d.Seats {
		if v == 2 {
			o++
		}
	}
	fmt.Println(o, " Occupied")
}

func (d *Day11) GetSeatEvent(s Seat, maxOccupied int) int {
	adjacentOccupied := 0
	adjacent, ok := d.AdjacentSeats[s]
	if ok {
		for k, _ := range adjacent {
			if d.Seats[k] == 2 {
				adjacentOccupied++
			}
		}
	} else {
		d.AdjacentSeats[s] = map[Seat]bool{}
		for i := s.I - 1; i <= s.I+1; i++ {
			for j := s.J - 1; j <= s.J+1; j++ {
				c := Seat{I: i, J: j}
				if s.I == i && s.J == j {
					continue
				}
				if d.Seats[c] == 2 { //occupied
					adjacentOccupied++
				}
				d.AdjacentSeats[s][c] = true
			}
		}
	}
	if d.Seats[s] == 1 && adjacentOccupied == 0 {
		return 2
	}
	if d.Seats[s] == 2 && adjacentOccupied >= maxOccupied {
		return 1
	}
	return -1
}

func (d *Day11) GetSeatEvent2(s Seat, maxOccupied int) int {
	adjacentOccupied := 0
	var c Seat
	if d.VisibleSeats[s] != nil {
		for k, _ := range d.VisibleSeats[s] {
			if d.Seats[k] == 2 {
				adjacentOccupied++
			}
		}
	} else {

		d.VisibleSeats[s] = map[Seat]bool{}
		for i := s.I - 1; i <= s.I+1; i++ {
			if i < 0 || i > d.Height-1 {
				continue
			}
			for j := s.J - 1; j <= s.J+1; j++ {
				if j < 0 || j > d.Width-1 {
					continue
				}
				c = Seat{I: i, J: j}
				if s.I == i && s.J == j {
					continue
				}
				for {
					v, exists := d.Seats[c]
					if !exists {
						break // edge
					}
					if v > 0 { // seat found
						break
					}
					if v == 0 { // floor, keep going
						//keep going in direction until you find a seat or edge
						c = Seat{I: c.I + (i - s.I), J: c.J + (j - s.J)}
					}
				}
				d.VisibleSeats[s][c] = true
				if d.Seats[c] == 2 { //occupied
					adjacentOccupied++
				}
			}
		}
	}
	if d.Seats[s] == 1 && adjacentOccupied == 0 {
		return 2
	}
	if d.Seats[s] == 2 && adjacentOccupied >= maxOccupied {
		return 1
	}
	return -1
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day11) Part2() {
	fmt.Println("Day 11 Part 2")
	d.VisibleSeats = map[Seat]map[Seat]bool{}
	c := 0
	for {
		seatsToOccupy := []Seat{}
		seatsToVacate := []Seat{}
		for i := 0; i < d.Height; i++ {
			for j := 0; j < d.Width; j++ {
				s := Seat{I: i, J: j}
				if d.Seats[s] == 0 {
					continue
				}
				switch d.GetSeatEvent2(s, 5) {
				case 1:
					seatsToVacate = append(seatsToVacate, s)
				case 2:
					seatsToOccupy = append(seatsToOccupy, s)
				}
			}
		}
		//	fmt.Println(c, len(seatsToOccupy), len(seatsToVacate))
		if len(seatsToOccupy) == 0 && len(seatsToVacate) == 0 {
			d.printGrid()
			d.CountOccupied()
			break
		}
		for _, s := range seatsToVacate {
			d.Seats[s] = 1
		}
		for _, s := range seatsToOccupy {
			d.Seats[s] = 2
		}

		c++
		if c%2 == 0 {
			d.printGrid()
			time.Sleep(100 * time.Millisecond)
		}

	}
}

func (d *Day11) printGrid() {
	colorGreen := "\033[32m"
	colorRed := "\033[31m"
	colorBlack := "\033[30m"
	c := ""
	var s int
	b := strings.Builder{}
	for i := 0; i < d.Height; i++ {
		for j := 0; j < d.Width; j++ {
			s = d.Seats[Seat{I: i, J: j}]
			switch s {
			case 0:
				c = colorBlack + " "
			case 1:
				c = colorRed + "█"
			case 2:
				c = colorGreen + "█"
			}
			b.WriteString(c)
		}
		b.WriteString("\n")
	}

	fmt.Printf("\033[H")
	fmt.Printf("%s", b.String())
}
