package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day01 struct{}

var _ Day = Day01{}

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
	diffs := d.Differences(left, right)

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

func (d Day01) Differences(left, right []int) []int {
	diffs := []int{}
	sort.Ints(left)
	sort.Ints(right)
	for i := range left {
		diffs = append(diffs, (int)(math.Abs((float64)(left[i]-right[i]))))
	}
	return diffs
}

// Part2 implements Day.
func (d Day01) Part2(input string) (int, error) {
	left, right, err := d.SplitLists(SplitInput(input))
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, i := range d.similarityScore(left, right) {
		sum += i
	}

	return sum, nil
}

func (d Day01) similarityScore(left, right []int) []int {
	counts := map[int]int{}
	for _, i := range right {
		counts[i] += 1
	}

	var scores []int
	for _, i := range left {
		scores = append(scores, i*counts[i])
	}
	return scores
}
