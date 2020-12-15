package day

import "time"

// Day is an interface for each day's solution to implement
type Day interface {
	ReadFile(path string) error
	Part1()
	Part2()
}

// GetToday calculates the current day of the month
// and returns the solution for that day
func GetToday() Day {
	_, _, d := time.Now().Date()
	return GetDay(d)
}

// GetDay takes a day of the month and returns the
// solution for that day
func GetDay(d int) Day {
	days := GetDays()
	return days[d-1]
}

// GetDays returns all the solutions in order
func GetDays() []Day {
	return []Day{
		&Day1{},
		&Day2{},
		&Day3{},
		&Day4{},
		&Day5{},
		&Day6{},
		&Day7{},
		&Day8{},
		&Day9{},
		&Day10{},
		&Day11{},
		&Day12{},
		&Day13{},
		&Day14{},
	}
}
