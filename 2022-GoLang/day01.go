package main

import (
	"strconv"
	"strings"
)

func CalorieCount(foodList string) int {
	var max = 0
	var sum = 0

	lines := strings.Split(foodList, "\n")
	for _, line := range lines {
		if line == "" {
			if sum > max {
				max = sum
			}

			sum = 0
		} else {
			num, _ := strconv.Atoi(line)
			sum += num
		}
	}
	return max
}

func CalorieCountThree(foodList string) int {
	var max, second, third, sum int

	lines := strings.Split(foodList, "\n")
	for _, line := range lines {
		if line == "" {
			if sum > max {
				max, second, third = sum, max, second
			} else if sum > second {
				second, third = sum, second
			} else if sum > third {
				third = sum
			}

			sum = 0
		} else {
			num, _ := strconv.Atoi(line)
			sum += num
		}
	}
	return max + second + third
}
