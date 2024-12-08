package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Day02 struct{}

var _ Day = Day02{}

// InputPath implements Day.
func (d Day02) InputPath() string {
	return "input/day02.txt"
}

// Part1 implements Day.
func (d Day02) Part1(in string) (int, error) {
	count := 0

	for _, line := range SplitInput(in) {
		safe, err := d.isSafe(line)

		if err != nil {
			return -1, nil
		}

		if safe {
			count++
		}
	}
	return count, nil
}

func (d Day02) isSafe(levelsStr string) (bool, error) {
	levels := strings.Split(levelsStr, " ")

	zero, err := strconv.Atoi(levels[0])
	if err != nil {
		return false, err
	}

	first, err := strconv.Atoi(levels[1])
	if err != nil {
		return false, err
	}

	decreasing := zero > first

	for i := 1; i < len(levels); i++ {
		last, err := strconv.Atoi(levels[i-1])
		if err != nil {
			return false, err
		}

		curr, err := strconv.Atoi(levels[i])
		if err != nil {
			return false, err
		}

		if decreasing {
			if last <= curr || (last-curr > 3) {
				return false, nil
			}
		} else {
			if last >= curr || (curr-last > 3) {
				return false, nil
			}
		}
	}

	return true, nil
}

// Part2 implements Day.
func (d Day02) Part2(string) (int, error) {
	return -1, fmt.Errorf("unimplemented")
}
