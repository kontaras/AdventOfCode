package main

import (
	"strconv"
	"strings"
)

func OverlappingAssignments(ranges string) int {
	lines := strings.Split(ranges, "\n")
	var overlap int
	for _, assignment := range lines {
		pair := strings.Split(assignment, ",")
		if len(pair) == 2 {
			min1, max1 := parseAssignment(pair[0])
			min2, max2 := parseAssignment(pair[1])

			if (min1 <= min2 && max1 >= max2) || (min1 >= min2 && max1 <= max2) {
				overlap++
			}
		}

	}

	return overlap
}

func parseAssignment(assignement string) (int, int) {
	pair := strings.Split(assignement, "-")

	min, err1 := strconv.Atoi(pair[0])
	if err1 != nil {
		panic(err1)
	}

	max, err2 := strconv.Atoi(pair[1])
	if err2 != nil {
		panic(err2)
	}

	return min, max
}
