package main

import (
	"log"
	"time"
	"vivanshah/aoc/day"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	start := time.Now()

	d := day.Day1{}
	d.Part1()
	d.Part2()

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
