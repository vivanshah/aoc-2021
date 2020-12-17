package day

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Day4 struct {
	Passports []Passport
}

type Passport struct {
	byr *int
	iyr *int
	eyr *int
	hgt string
	hcl string
	ecl string
	pid string
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day4) ReadFile(path string) error {
	fmt.Println("Reading input")
	file, err := os.Open("../../day4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string
	var pp Passport
	pp = Passport{}
	for {
		line, err = reader.ReadString('\n')

		if err != nil && err != io.EOF {
			d.Passports = append(d.Passports, pp)
			pp = Passport{}
			break
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			d.Passports = append(d.Passports, pp)
			pp = Passport{}
			continue
		}
		fields := strings.Split(line, " ")

		for _, f := range fields {
			parts := strings.Split(f, ":")
			switch parts[0] {
			case "byr":
				var i int
				i, _ = strconv.Atoi(parts[1])
				pp.byr = &i
			case "iyr":
				var i int
				i, _ = strconv.Atoi(parts[1])
				pp.iyr = &i
			case "eyr":
				var i int
				i, _ = strconv.Atoi(parts[1])
				pp.eyr = &i
			case "hgt":
				pp.hgt = parts[1]
			case "hcl":
				pp.hcl = parts[1]
			case "ecl":
				pp.ecl = parts[1]
			case "pid":
				pp.pid = parts[1]
			}
		}

		if err != nil {
			d.Passports = append(d.Passports, pp)
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}
	fmt.Println("Read ", len(d.Passports), " passports")
	return nil
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day4) Part1() {
	fmt.Println("Day 4 Part 1")
	valid := 0
	for _, p := range d.Passports {
		if p.byr != nil && p.iyr != nil && p.hgt != "" && p.eyr != nil && p.hcl != "" && p.ecl != "" && p.pid != "" {

			valid++
		}
	}
	fmt.Println(valid, " valid passports")
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day4) Part2() {
	fmt.Println("Day 4 Part 2")
	valid := 0
	for _, p := range d.Passports {
		if p.byr != nil && p.iyr != nil && p.hgt != "" && p.eyr != nil && p.hcl != "" && p.ecl != "" && p.pid != "" {
			if !(*p.byr >= 1920 && *p.byr <= 2020) {
				//fmt.Println("invalid birth year ", *p.byr)
				continue
			}
			if !(*p.iyr >= 2010 && *p.iyr <= 2020) {
				//fmt.Println("invalid issue year ", *p.iyr)
				continue
			}
			if !(*p.eyr >= 2020 && *p.eyr <= 2030) {
				//	fmt.Println("invalid expiry year ", *p.eyr)
				continue
			}
			h, _ := regexp.MatchString("^[0-9]*(cm|in)", p.hgt)
			if !h {
				//fmt.Println("invalid height string", p.hgt)
				continue
			}
			if strings.HasSuffix(p.hgt, "cm") {
				var h int
				fmt.Sscanf(p.hgt, "%dcm", &h)
				if !(h >= 150 && h <= 193) {
					//	fmt.Println("invalid height: ", p.hgt)
					continue
				}
			} else if strings.HasSuffix(p.hgt, "in") {
				var h int
				fmt.Sscanf(p.hgt, "%din", &h)
				if !(h >= 59 && h <= 76) {
					//	fmt.Println("invalid height: ", p.hgt)
					continue
				}
			}
			if rune(p.hcl[0]) != rune('#') {
				continue
			}
			m, _ := regexp.MatchString("^#[a-f0-9]{6}$", p.hcl)
			if !m {
				continue
			}

			e, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", p.ecl)
			if !e {
				continue
			}

			pid, _ := regexp.MatchString("^[0-9]{9}$", p.pid)
			if !pid {
				continue
			}
			valid++
		}
	}
	fmt.Println(valid, " valid passports")
}
