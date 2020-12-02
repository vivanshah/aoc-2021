package day

// Day is an interface for each day's solution to implement
type Day interface {
	ReadFile(path string) ([]string, error)
	Part1()
	Part2()
}
