package main

import (
	"fmt"
	"os"
)

func main() {
	day1()
	day2()
	day3()
}

func day1() {
	file, err := os.ReadFile("input/day01.txt")

	if err != nil {
		panic("Could not read input file: " + err.Error())
	}

	var value = CalorieCount(string(file))
	fmt.Printf("Day 1 Part 1: %d\n", value)

	value = CalorieCountThree(string(file))
	fmt.Printf("Day 1 Part 2: %d\n", value)
}

func day2() {
	file, err := os.ReadFile("input/day02.txt")

	if err != nil {
		panic("Could not read input file: " + err.Error())
	}

	var value = RpsScore(string(file))
	fmt.Printf("Day 2 Part 1: %d\n", value)

	value = RpsPlays(string(file))
	fmt.Printf("Day 2 Part 2: %d\n", value)
}

func day3() {
	file, err := os.ReadFile("input/day03.txt")

	if err != nil {
		panic("Could not read input file: " + err.Error())
	}

	var value = RucksackPriorityScore(string(file))
	fmt.Printf("Day 3 Part 1: %d\n", value)

	value = GroupBadge(string(file))
	fmt.Printf("Day 3 Part 2: %d\n", value)
}
