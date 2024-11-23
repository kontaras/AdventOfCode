package main

import (
	"math"
	"strconv"
	"strings"
)

func Day03Part1(input string) int {
	sum := 0

	for _, num := range GetNumbers(input) {
		sum += num
	}

	return sum
}

func GetNumbers(input string) []int {
	numbers := []int{}

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		number := []byte{}
		j := 0
		start := -1
		for j < len(line) {
			char := line[j]
			if char >= '0' && char <= '9' {
				if start < 0 {
					start = j
				}
				number = append(number, char)
			} else if len(number) > 0 {
				if CheckNumber(lines, i, start, j-1) {
					num, err := strconv.Atoi(string(number))
					if err != nil {
						panic(err)
					}
					numbers = append(numbers, num)
				}
				number = []byte{}
				start = -1
			}
			j++
		}

		if len(number) > 0 {
			if CheckNumber(lines, i, start, j-1) {
				num, err := strconv.Atoi(string(number))
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, num)
			}
		}
	}

	return numbers
}

func CheckNumber(lines []string, lineNum, start, end int) bool {
	isSymbol := func(char byte) bool {
		return char != '.' && char != '\n'
	}

	if lineNum > 0 {
		line := lines[lineNum-1]
		for i := math.Max(0, float64(start-1)); i < math.Min(float64(end+2), float64(len(line))); i++ {
			if isSymbol(line[int(i)]) {
				return true
			}
		}
	}

	line := lines[lineNum]
	if start > 0 && isSymbol(line[start-1]) {
		return true
	}

	if end+1 < len(line) && isSymbol(line[end+1]) {
		return true
	}

	if lineNum+1 < len(lines) {
		line := lines[lineNum+1]
		for i := math.Max(0, float64(start-1)); i < math.Min(float64(end+2), float64(len(line))); i++ {
			if isSymbol(line[int(i)]) {
				return true
			}
		}
	}

	return false
}
