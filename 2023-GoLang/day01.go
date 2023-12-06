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
