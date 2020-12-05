package main

import (
	"flag"
	"fmt"
	"time"
	"vivanshah/aoc/day"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dayToRun := flag.Int("day", -1, "Specify which day to run")
	flag.Parse()
	fmt.Println(dayToRun)
	var days []day.Day
	if dayToRun != nil && *dayToRun != -1 {
		days = []day.Day{day.GetDay(*dayToRun)}
	} else {
		days = day.GetDays()
	}
	fmt.Println("Running ", len(days), " days")
	for i, d := range days {
		start := time.Now()
		d.ReadFile("../../day" + fmt.Sprint(i+1) + ".txt")
		d.Part1()
		elapsed := time.Since(start)
		fmt.Printf("Part 1 took %s\n", elapsed)
		start = time.Now()
		d.Part2()
		elapsed = time.Since(start)
		fmt.Printf("Part 2 took %s\n", elapsed)
	}
}
