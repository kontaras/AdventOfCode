package main

import (
	"fmt"
	"os"
)

func main() {
	day1()
	day2()
	day3()
	day4()
	day5()
	day6()
	day07()
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

func day4() {
	file, err := os.ReadFile("input/day04.txt")

	if err != nil {
		panic("Could not read input file: " + err.Error())
	}

	var value = SubsetAssignments(string(file))
	fmt.Printf("Day 4 Part 1: %d\n", value)

	value = OverlappingAssignments(string(file))
	fmt.Printf("Day 4 Part 2: %d\n", value)
}

func day5() {
	file, err := os.ReadFile("input/day05.txt")

	if err != nil {
		panic("Could not read input file: " + err.Error())
	}

	var value = CrateStack(string(file))
	fmt.Printf("Day 5 Part 1: '%s'\n", value)

	value = CrateStack2(string(file))
	fmt.Printf("Day 5 Part 2: '%s'\n", value)
}

func day6() {
	file, err := os.ReadFile("input/day06.txt")

	if err != nil {
		panic("Could not read input file: " + err.Error())
	}

	var value = FindPacketStart(string(file))
	fmt.Printf("Day 6 Part 1: %d\n", value)

	value = FindMessageStart(string(file))
	fmt.Printf("Day 6 Part 2: %d\n", value)
}

func day07() {
	file, err := os.ReadFile("input/day07.txt")

	if err != nil {
		panic("Could not read input file: " + err.Error())
	}

	var value = DirsToDelete(string(file))
	fmt.Printf("Day 7 Part 1: %d\n", value)

	value = FindOptimalDelete(string(file))
	fmt.Printf("Day 7 Part 2: %d\n", value)
}
