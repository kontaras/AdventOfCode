package main

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"maps"
)

func Day03Part1(input string) int {
	sum := 0

	for _, num := range GetNumbers(input) {
		sum += num
	}

	return sum
}

func Day03Part2(input string) int {
	sum := 0

	for _, num := range GetGearRatios(input) {
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
				if checkNumber(lines, i, start, j-1) {
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
			if checkNumber(lines, i, start, j-1) {
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

func checkNumber(lines []string, lineNum, start, end int) bool {
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

func GetGearRatios(input string) []int {
	numbers := []int{}
	for _, gear := range GetGears(input) {
		if len(gear) == 2 {
			numbers = append(numbers, gear[0]*gear[1])
		}
	}

	return numbers
}

func GetGears(input string) [][]int {
	gears := map[gear][]int{}

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
				if thisGear, found := getGears(lines, i, start, j-1); found {
					num, err := strconv.Atoi(string(number))
					if err != nil {
						panic(err)
					}
					if _, ok := gears[thisGear]; !ok {
						gears[thisGear] = []int{}
					}
					gears[thisGear] = append(gears[thisGear], num)
				}
				number = []byte{}
				start = -1
			}
			j++
		}

		if len(number) > 0 {
			if thisGear, found := getGears(lines, i, start, j-1); found {
				num, err := strconv.Atoi(string(number))
				if err != nil {
					panic(err)
				}
				if _, ok := gears[thisGear]; !ok {
					gears[thisGear] = []int{}
				}
				gears[thisGear] = append(gears[thisGear], num)
			}
		}
	}

	return slices.Collect(maps.Values(gears))
}

type gear struct {
	x int
	y int
}

func getGears(lines []string, lineNum, start, end int) (gear, bool) {
	isSymbol := func(char byte) bool {
		return char == '*'
	}

	if lineNum > 0 {
		line := lines[lineNum-1]
		for i := math.Max(0, float64(start-1)); i < math.Min(float64(end+2), float64(len(line))); i++ {
			if isSymbol(line[int(i)]) {
				return gear{lineNum - 1, int(i)}, true
			}
		}
	}

	line := lines[lineNum]
	if start > 0 && isSymbol(line[start-1]) {
		return gear{lineNum, start - 1}, true
	}

	if end+1 < len(line) && isSymbol(line[end+1]) {
		return gear{lineNum, end + 1}, true
	}

	if lineNum+1 < len(lines) {
		line := lines[lineNum+1]
		for i := math.Max(0, float64(start-1)); i < math.Min(float64(end+2), float64(len(line))); i++ {
			if isSymbol(line[int(i)]) {
				return gear{lineNum + 1, int(i)}, true
			}
		}
	}

	return gear{}, false
}
