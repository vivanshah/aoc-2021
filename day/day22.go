package day

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Day22 struct {
	MyDeck   *list.List
	CrabDeck *list.List
	Total    int
}

// ReadFile reads a file
func (d *Day22) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day22.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	section := 0
	d.MyDeck = list.New()
	d.CrabDeck = list.New()
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			section++
			continue
		}
		if strings.HasPrefix(l, "Player") {
			continue
		}
		if section == 0 {
			d.MyDeck.PushBack(GetInt(l))
		}
		if section == 1 {
			d.CrabDeck.PushBack(GetInt(l))
		}
		d.Total++
	}

	fmt.Println("Read ", d.CrabDeck.Len()+d.MyDeck.Len(), " Cards")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day22) Part1() {
	fmt.Println("Day 22 Part 1")
	rounds := 1
	for d.MyDeck.Len() != d.Total && d.CrabDeck.Len() != d.Total {
		m := d.MyDeck.Remove(d.MyDeck.Front())
		c := d.CrabDeck.Remove(d.CrabDeck.Front())
		myTop := m.(int)
		crabTop := c.(int)
		fmt.Printf("Mine: %d Crabs: %d\n", myTop, crabTop)
		if myTop > crabTop {
			d.MyDeck.PushBack(myTop)
			d.MyDeck.PushBack(crabTop)
		} else {
			d.CrabDeck.PushBack(crabTop)
			d.CrabDeck.PushBack(myTop)
		}
		//fmt.Println(rounds)
		rounds++
	}
	var winner *list.List

	if d.MyDeck.Len() == d.Total {
		winner = d.MyDeck
	} else {
		winner = d.CrabDeck
	}
	//i win!
	total := 0

	for x := d.Total; x > 0; x-- {
		c := winner.Remove(winner.Front())
		card := c.(int)
		total = total + (x * card)
	}
	fmt.Println(total)

}

// Part2 executes part 2 of of this day's puzzle
func (d *Day22) Part2() {
	fmt.Println("Day 22 Part 2")
	input, _ := ioutil.ReadFile("../../day22.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\r\n\r\n")

	decks := make([][]int, len(split))
	for i, s := range split {
		for _, s := range strings.Split(s, "\r\n")[1:] {
			c, _ := strconv.Atoi(s)
			decks[i] = append(decks[i], c)
		}
	}

	_, score := run([][]int{append([]int{}, decks[0]...), append([]int{}, decks[1]...)}, false)
	fmt.Println(score)
	_, score = run(decks, true)
	fmt.Println(score)
}

func run(ds [][]int, rec bool) (win int, score int) {
	seen := map[string]struct{}{}

	for len(ds[0]) > 0 && len(ds[1]) > 0 {
		win = 0
		if _, ok := seen[fmt.Sprint(ds)]; rec && ok {
			break
		}
		seen[fmt.Sprint(ds)] = struct{}{}

		if rec && len(ds[0]) > ds[0][0] && len(ds[1]) > ds[1][0] {
			win, _ = run([][]int{append([]int{}, ds[0][1:ds[0][0]+1]...), append([]int{}, ds[1][1:ds[1][0]+1]...)}, rec)
		} else if ds[0][0] < ds[1][0] {
			win = 1
		}

		ds[win] = append(ds[win], ds[win][0], ds[-win+1][0])
		ds[0], ds[1] = ds[0][1:], ds[1][1:]
	}

	for i, c := range ds[win] {
		score += c * (len(ds[win]) - i)
	}
	return
}
