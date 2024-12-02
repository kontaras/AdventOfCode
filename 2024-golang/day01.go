package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day01 struct{}

// InputPath implements Day.
func (d Day01) InputPath() string {
	return "input/day01.txt"
}

// Part1 implements Day.
func (d Day01) Part1(input string) (int, error) {
	left, right, err := d.SplitLists(SplitInput(input))
	if err != nil {
		return 0, err
	}
	diffs := d.Diffences(left, right)

	sum := 0
	for _, num := range diffs {
		sum += num
	}

	return sum, nil
}

func (d Day01) SplitLists(lines []string) ([]int, []int, error) {
	left := []int{}
	right := []int{}

	for _, line := range lines {
		numbers := strings.SplitN(line, "   ", 2)
		first, err := strconv.Atoi(numbers[0])
		if err != nil {
			return left, right, err
		}
		left = append(left, first)

		second, err := strconv.Atoi(numbers[1])
		if err != nil {
			return left, right, err
		}
		right = append(right, second)
	}
	return left, right, nil
}

func (d Day01) Diffences(left, right []int) []int {
	diffs := []int{}
	sort.Ints(left)
	sort.Ints(right)
	for i := range left {
		diffs = append(diffs, (int)(math.Abs((float64)(left[i]-right[i]))))
	}
	return diffs
}

// Part2 implements Day.
func (d Day01) Part2(string) (int, error) {
	return 0, fmt.Errorf("Unimplemented")
}

var _ Day = Day01{}
