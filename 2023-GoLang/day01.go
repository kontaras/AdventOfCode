package main

import (
	"strconv"
	"strings"
)

func Day01Part1(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sum += CalibrateLine(line)
	}

	return sum
}

func CalibrateLine(line string) int {
	if len(line) < 1 {
		// Handle empty lines
		return 0
	}

	first := byte(0)
	var last byte
	for i := 0; i < len(line); i++ {
		char := line[i]
		if char >= '0' && char <= '9' {
			if first == 0 {
				first = char
			}
			last = char
		}
	}
	answer, err := strconv.Atoi(string([]byte{first, last}))
	if err != nil {
		panic("Could not parse number " + err.Error())
	}
	return answer
}

func Day01Part2(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sum += CalibrateLineWords(line)
	}

	return sum
}

func CalibrateLineWords(line string) int {
	if len(line) < 1 {
		// Handle empty lines
		return 0
	}

	first := -1
	var last int
	for i := 0; i < len(line); i++ {
		char := ParseNumber(line[i:])
		if char >= 0 {
			if first == -1 {
				first = char
			}
			last = char
		}
	}

	return first*10 + last
}

var LOOKUP_TABLE = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func ParseNumber(component string) int {
	for str, val := range LOOKUP_TABLE {
		if strings.HasPrefix(component, str) {
			return val
		}
	}
	return -1
}
