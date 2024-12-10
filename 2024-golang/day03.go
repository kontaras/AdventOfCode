package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Day03 struct{}

var _ Day = Day03{}

// InputPath implements Day.
func (d Day03) InputPath() string {
	return "input/day03.txt"
}

// Part1 implements Day.
func (d Day03) Part1(input string) (int, error) {
	lines := SplitInput(input)
	sum := 0
	for _, line := range lines {
		muls, err := d.findMuls(line)
		if err != nil {
			return -1, err
		}

		for _, mul := range muls {
			sum += mul.a * mul.b
		}
	}

	return sum, nil
}

type mul struct {
	a, b int
}

func (d Day03) findMuls(memory string) ([]mul, error) {
	muls := []mul{}
	re, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		return muls, err
	}

	matches := re.FindAllStringSubmatch(memory, -1)

	for _, match := range matches {
		first, err := strconv.Atoi(match[1])
		if err != nil {
			return muls, err
		}
		second, err := strconv.Atoi(match[2])
		if err != nil {
			return muls, err
		}
		muls = append(muls, mul{first, second})
	}

	return muls, nil
}

// Part2 implements Day.
func (d Day03) Part2(input string) (int, error) {
	lines := SplitInput(input)
	line := strings.Join(lines, "")
	sum := 0
	muls, err := d.findEnabledMuls(line)
	if err != nil {
		return -1, err
	}

	for _, mul := range muls {
		sum += mul.a * mul.b
	}

	return sum, nil
}

func (d Day03) findEnabledMuls(memory string) ([]mul, error) {
	muls := []mul{}
	re, err := regexp.Compile(`(mul\((\d{1,3}),(\d{1,3})\))|(do(n't)?\(\))`)
	//re, err := regexp.Compile(`do`)
	if err != nil {
		return muls, err
	}

	matches := re.FindAllStringSubmatch(memory, -1)

	do := true
	for _, match := range matches {
		switch match[4] {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if do {
				first, err := strconv.Atoi(match[2])
				if err != nil {
					return muls, err
				}
				second, err := strconv.Atoi(match[3])
				if err != nil {
					return muls, err
				}
				muls = append(muls, mul{first, second})
			}
		}

	}

	return muls, nil
}
