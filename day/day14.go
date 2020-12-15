package day

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Day14 struct {
	Instruction []string
}

// ReadFile reads a file and returns a slice of strings, one for
// each line
func (d *Day14) ReadFile(path string) error {
	fmt.Println("Reading input")
	input, err := ioutil.ReadFile("../../day14.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\r\n")
	d.Instruction = []string{}
	for _, l := range lines {
		l = strings.TrimSpace(l)
		d.Instruction = append(d.Instruction, l)
	}
	fmt.Println("Read ", len(lines), " Instructions")
	return nil
}

func GetMask(s string) string {
	return s[7:]
}

func GetMem(s string) (int, uint64) {
	var i int
	var v uint64
	parts := strings.Split(s, " = ")
	m := strings.ReplaceAll(parts[0][:len(parts[0])-1], "mem[", "")
	v, _ = strconv.ParseUint(parts[1], 10, 64)
	i, _ = strconv.Atoi(m)
	return i, v
}

// Part1 executes part 1 of of this day's puzzle
func (d *Day14) Part1() {
	fmt.Println("Day 14 Part 1")
	mem := map[int]uint64{}
	var mask string
	for _, i := range d.Instruction {
		if strings.HasPrefix(i, "mask") {
			mask = GetMask(i)
		} else {
			m, val := GetMem(i)
			//fmt.Printf("val: %036b\n", val)
			a := strings.Replace(mask, "X", "0", -1)
			//fmt.Printf("a  : %s\n", a)
			aInt, _ := strconv.ParseUint(a, 2, 64)
			val = val | aInt
			//fmt.Printf("val: %036b\n", val)
			b := strings.Replace(mask, "X", "1", -1)
			//fmt.Printf("b  : %s\n", b)
			bInt, _ := strconv.ParseUint(b, 2, 64)
			val = val & bInt
			//	fmt.Printf("val: %036b\n", val)
			//fmt.Printf("mem[%d] = %d\n", m, val)
			mem[m] = val
		}
	}
	var s uint64
	for _, v := range mem {
		s += v
	}
	fmt.Println(s)
}

// Part2 executes part 2 of of this day's puzzle
func (d *Day14) Part2() {
	fmt.Println("Day 14 Part 2")
	mem := map[uint64]uint64{}
	var mask string
	for _, i := range d.Instruction {
		if strings.HasPrefix(i, "mask") {
			mask = GetMask(i)
		} else {
			//fmt.Println("mask: ", mask)
			m, val := GetMem(i)
			uM := uint64(m)
			//fmt.Printf("uM   : %036b\n", uM)
			a := strings.Replace(mask, "X", "0", -1)
			aInt, _ := strconv.ParseUint(a, 2, 64)
			uM = uM | aInt
			//fmt.Printf("uM   : %036b\n", uM)
			b := strings.Replace(mask, "X", "0", -1)
			bInt, _ := strconv.ParseUint(b, 2, 64)
			uM = uM | bInt
			//fmt.Printf("uM   : %036b\n", uM)
			m2 := fmt.Sprintf("%036b", uM)
			var initial string
			for i, j := range mask {
				if j == rune('X') {
					m2 = replaceAtIndex(m2, 'X', i)
				}
			}
			initial = strings.ReplaceAll(m2, "X", "0")
			uM, _ = strconv.ParseUint(initial, 2, 64)
			//fmt.Println(m2)
			addresses := GetMemAddresses(m2, uM)
			//fmt.Println(addresses)
			for _, a := range addresses {
				mem[a] = val
			}
		}
	}
	var s uint64
	for _, v := range mem {
		s += v
	}
	fmt.Println(s)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
func GetMemAddresses(mask string, m uint64) []uint64 {
	r := []rune(mask)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	reverse := string(r)
	//fmt.Printf("m    : %036b\n", m)
	//fmt.Println(mask)
	//fmt.Println(reverse)

	dp := map[uint64][]uint64{}
	dp[0] = []uint64{m}
	var i uint64
	i = 0
	for {
		if i > 35 {
			break
		}
		//fmt.Println("considering first ", i+1, " bits of mask ", string(reverse[i]))
		if dp[i+1] == nil {
			dp[i+1] = []uint64{}
		}
		//fmt.Println(dp)
		for _, number := range dp[i] {
			switch string(reverse[i]) {
			case "0":
				x := (1 << i) & m
				dp[i+1] = append(dp[i+1], number|x)
			//	fmt.Println("appended ", number|x)

			case "1":
				x := uint64(1 << i)
				dp[i+1] = append(dp[i+1], number|x)
				//fmt.Println("appended ", number|x)
			case "X":
				x := uint64(1 << i)
				//fmt.Println(x)
				//fmt.Println("appended ", number|x)
				dp[i+1] = append(dp[i+1], number|x)
				//fmt.Println("appended ", number|(x&m))
				dp[i+1] = append(dp[i+1], number|(x&m))
			default:
				fmt.Printf("Unexpected mask value >%v<", string(reverse[i]))
			}
		}
		//fmt.Println(i, dp[i+1])
		i++
	}
	result := []uint64{}
	e := map[uint64]bool{}
	for x, v := range dp {
		if x == 0 {
			continue
		}
		for _, m := range v {
			if !e[m] {
				result = append(result, m)
			}

			e[m] = true
		}
	}
	return result
}
