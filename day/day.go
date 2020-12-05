package day

import "time"

// Day is an interface for each day's solution to implement
type Day interface {
	ReadFile(path string) error
	Part1()
	Part2()
}

func GetToday() Day {
	_, _, d := time.Now().Date()
	return GetDay(d)
}

func GetDay(d int) Day {
	days := GetDays()
	return days[d-1]
}

func GetDays() []Day {
	return []Day{
		&Day1{},
		&Day2{},
		&Day3{},
		&Day4{},
		&Day5{},
	}
}
