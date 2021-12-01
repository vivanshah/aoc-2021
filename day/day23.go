package day

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Day23 struct {
	Cups  *list.List
	Min   int
	Max   int
	Total int
}

// ReadFile reads a file
func (d *Day23) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day23.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	d.Cups = list.New()
	d.Min = math.MaxInt64
	d.Max = math.MinInt64
	for _, l := range lines {
		l = strings.TrimSpace(l)
		for _, c := range l {
			cup := GetInt(string(c))
			d.Cups.PushBack(cup)
			d.Total++
			if cup < d.Min {
				d.Min = cup
			}
			if cup > d.Max {
				d.Max = cup
			}
		}
	}
	d.Cups.Back()

	fmt.Println("Read ", d.Cups.Len(), " Cups")
	fmt.Println(d.Min, d.Max)

	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day23) Part1() {
	fmt.Println("Day 23 Part 1")
	rounds := 100

	current := d.Cups.Front()
	for round := 0; round < rounds; round++ {

		r := current.Next()
		if r == nil {
			r = d.Cups.Front()
		}

		pickedUpMap := map[int]*list.Element{}
		pickedUp := []*list.Element{}
		for i := 0; i < 3; i++ {
			pickedUpMap[r.Value.(int)] = r
			pickedUp = append(pickedUp, r)
			r = r.Next()
			if r == nil {
				r = d.Cups.Front()
			}
		}
		var dest *list.Element
		target := current.Value.(int) - 1
		if target < d.Min {
			target = d.Max
		}
		for {
			if pickedUpMap[target] == nil {
				break
			}
			target--
			if target < d.Min {
				target = d.Max
			}
		}

		i := current.Next()
		if i == nil {
			i = d.Cups.Front()
		}
		for dest == nil {
			if i.Value.(int) == target {
				dest = i
				break
			}
			i = i.Next()
			if i == nil {
				i = d.Cups.Front()
			}
		}
		for _, f := range pickedUp {
			d.Cups.MoveAfter(f, dest)
			dest = dest.Next()
		}
		current = current.Next()
		if current == nil {
			current = d.Cups.Front()
		}
	}

	p := d.Cups.Front()
	for p.Value.(int) != 1 {
		p = p.Next()
	}
	p = p.Next()
	for x := 0; x < d.Cups.Len()-1; x++ {
		fmt.Printf("%d", p.Value.(int))
		p = p.Next()
		if p == nil {
			p = d.Cups.Front()
		}
	}
	fmt.Println()

}

func (d *Day23) PrintCups() {
	fmt.Print("\n")

	p := d.Cups.Front()
	for x := 0; x < d.Cups.Len(); x++ {
		fmt.Print(p.Value.(int))
		p = p.Next()
	}
	fmt.Print("\n")
}
func (d *Day23) GetState() state {

	place := 1
	result := 0
	p := d.Cups.Back()
	for x := 0; x < d.Cups.Len(); x++ {
		v := p.Value.(int)
		result = result + (place * v)
		place = place * 10
		p = p.Prev()
	}
	return state{State: result}
}

func (d *Day23) Find(t int) *list.Element {
	//fmt.Println("len:", d.Cups.Len())
	p := d.Cups.Back()
	for x := 0; x < d.Cups.Len(); x++ {
		v := p.Value.(int)
		//	fmt.Println(v)
		if v == t {
			return p
		}
		p = p.Prev()
	}
	panic(fmt.Sprintf("Couldn't find %d", t))
}

type state struct {
	Current int
	State   int
}

func (d *Day23) SetState(s state) {
	d.Cups = list.New()
	for i := 9; i > 0; i-- {
		d.Cups.PushFront(digit(s.State, i))
	}
}
func digit(num, place int) int {
	r := num % int(math.Pow(10, float64(place)))
	return r / int(math.Pow(10, float64(place-1)))
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day23) Part2() {
	fmt.Println("Day 23 Part 2")

	for i := d.Max + 1; i <= 1000000; i++ {
		d.Cups.PushBack(i)
	}

	rounds := 10000000
	nextStates := map[state]state{}

	current := d.Cups.Front()
	var state state
	for round := 0; round < rounds; round++ {

		//if round%1000 == 0 {
		fmt.Println("Round:", round)
		//}
		state = d.GetState()
		state.Current = current.Value.(int)
		//fmt.Println(state)
		nextState, ok := nextStates[state]
		if ok {
			fmt.Println("previous state found at round ", round)

			//fmt.Println(nextState)
			d.SetState(nextState)
			current = d.Find(nextState.Current)
			continue
		}

		r := current.Next()
		if r == nil {
			r = d.Cups.Front()
		}

		pickedUpMap := map[int]*list.Element{}
		pickedUp := []*list.Element{}
		for i := 0; i < 3; i++ {
			pickedUpMap[r.Value.(int)] = r
			pickedUp = append(pickedUp, r)
			r = r.Next()
			if r == nil {
				r = d.Cups.Front()
			}
		}
		var dest *list.Element
		target := current.Value.(int) - 1
		if target < d.Min {
			target = d.Max
		}
		for {
			if pickedUpMap[target] == nil {
				break
			}
			target--
			if target < d.Min {
				target = d.Max
			}
		}

		i := current.Next()
		if i == nil {
			i = d.Cups.Front()
		}
		for dest == nil {
			if i.Value.(int) == target {
				dest = i
				break
			}
			i = i.Next()
			if i == nil {
				i = d.Cups.Front()
			}
		}
		for _, f := range pickedUp {
			d.Cups.MoveAfter(f, dest)
			dest = dest.Next()
		}
		current = current.Next()
		if current == nil {
			current = d.Cups.Front()
		}
		next := d.GetState()
		next.Current = current.Value.(int)
		nextStates[state] = next
	}

	p := d.Cups.Front()
	for p.Value.(int) != 1 {
		p = p.Next()
	}
	p = p.Next()
	for x := 0; x < d.Cups.Len()-1; x++ {
		fmt.Printf("%d", p.Value.(int))
		p = p.Next()
		if p == nil {
			p = d.Cups.Front()
		}
	}
	fmt.Println()

	p = d.Cups.Front()
	for p.Value.(int) != 1 {
		p = p.Next()
	}
	p = p.Next()
	for x := 0; x < 2; x++ {
		fmt.Printf("%d\t", p.Value.(int))
		p = p.Next()
		if p == nil {
			p = d.Cups.Front()
		}
	}
	fmt.Println()
}
