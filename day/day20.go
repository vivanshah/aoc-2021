package day

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Day20 struct {
	Tiles []tile
}

type tile struct {
	id     int
	data   map[coordinate]bool
	height int
	width  int
}

func (t tile) GetRotated(d int) tile {
	//	r := tile{id: t.id, height: t.width, width: t.height, data: map[coordinate]bool{}}
	return tile{}
}

func (t tile) GetVflip() tile {
	return tile{}
}
func (t tile) GetRow(i int) string {
	b := strings.Builder{}
	var d bool
	for j := 0; j < t.width; j++ {
		d = t.data[coordinate{I: i, J: j}]
		if d {
			b.WriteString("#")
		} else {
			b.WriteString(".")
		}
	}
	return b.String()
}
func (t tile) GetColumn(j int) string {
	b := strings.Builder{}
	var d bool
	for i := 0; i < t.height; i++ {
		d = t.data[coordinate{I: i, J: j}]
		if d {
			b.WriteString("#")
		} else {
			b.WriteString(".")
		}
	}
	return b.String()
}

// ReadFile reads a file
func (d *Day20) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day20.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	d.Tiles = []tile{}
	var cT tile
	tileRow := 0
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if len(l) == 0 { //break between tiles
			d.Tiles = append(d.Tiles, cT)
			continue
		}
		if strings.HasPrefix(l, "Tile") {
			// start of new tile
			p := strings.Fields(l)
			id, _ := strconv.Atoi(strings.ReplaceAll(p[1], ":", ""))
			cT = tile{id: id, height: 10, width: 10, data: map[coordinate]bool{}}
			tileRow = 0
			continue
		}

		//in a tile
		for j, c := range l {
			b := string(c) == "#"
			cT.data[coordinate{I: tileRow, J: j}] = b
		}
		tileRow++

	}

	fmt.Println("Read ", len(d.Tiles), " Tiles")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day20) Part1() {
	fmt.Println("Day 20 Part 1")
	edges := map[string]int{}
	corners := map[int]bool{}
	for _, t := range d.Tiles {
		fmt.Println("Tile: ", t.id)
		fR := t.GetRow(0)
		fmt.Println(fR)
		edges[fR]++
		lR := t.GetRow(t.height - 1)
		fmt.Println(lR)
		edges[lR]++
		fC := t.GetColumn(0)
		fmt.Println(fC)
		edges[fC]++
		lC := t.GetColumn(t.width - 1)
		fmt.Println(lC)
		edges[lC]++
	}
	for _, t := range d.Tiles {

		fR := t.GetRow(0)
		lR := t.GetRow(t.height - 1)
		fC := t.GetColumn(0)
		lC := t.GetColumn(t.width - 1)
		s := []string{fR, lR, fC, lC}
		m := 0
		for _, e := range s {
			if edges[e] == 1 {
				m++
			}
		}
		if m == 2 {
			corners[t.id] = true
		}

	}
	fmt.Println("scanning edges")
	var m int64
	m = 1
	for k := range corners {
		fmt.Println(k)
		m = m * int64(k)
	}
	fmt.Println(m)

}

// Part2 executes part 2 of of this day's puzzle
func (d *Day20) Part2() {
	fmt.Println("Day 20 Part 2")

}

/*
Read in each tile creating a dictionary of points that contain '#'
for each tile, rotate 90, 180 and 270 degrees, and also flip vertical and rotate 90, 180, 270 to get all permutations of flipped and rotated for each tile
For each tile create a list of other tiles that has a bottom row that matches the top row of current tile (not included flipped or rotated of itself)
For each tile create a list of other tiles that has a right row that matches the left row of current tile (not included flipped or rotated of itself)
Brute force the images by starting in top left corner building up an image where the top and left corners match of a tile
This returns all the images, rotated and flipped, so taking the first one and getting the corners gave me the answer.

Part 1 and then for each image
create a new image by combining the tiles , removing the edges and shifting the coords depending on location of the tile. - this is pretty grim, lots of off by 1 things going on and the code is a mess.
At each '#' check if there is a sea monster.*/
