package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
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
	all := flag.Bool("all", false, "Run All Days")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	var days []day.Day
	if dayToRun != nil && *dayToRun != -1 {
		days = []day.Day{day.GetDay(*dayToRun)}
	} else if all != nil && *all {
		days = day.GetDays()
	} else {
		days = []day.Day{day.GetToday()}
	}
	fmt.Println("Running ", len(days), " days")
	zero := time.Now()
	for i, d := range days {
		start := time.Now()
		d.ReadFile("../../day" + fmt.Sprint(i+1) + ".txt")
		d.Part1()
		elapsed := time.Since(start)
		fmt.Printf("Part 1 took %s\n", elapsed)
		d.ReadFile("../../day" + fmt.Sprint(i+1) + ".txt")
		start = time.Now()
		d.Part2()
		elapsed = time.Since(start)
		fmt.Printf("Part 2 took %s\n", elapsed)
	}

	fmt.Printf("Total Time Elapsed: %s\n", time.Since(zero))
}
