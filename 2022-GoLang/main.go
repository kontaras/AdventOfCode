package main

import (
	"fmt"
	"os"
)

func main() {
	day1()
	day2()
}

func day1() {
	file, err := os.ReadFile("input/day01.txt")

	if err != nil {
		panic("Could not read input file " + err.Error())
	}

	value := CalorieCount(string(file))
	fmt.Printf("Day 1 Part 1: %d\n", value)

	value2 := CalorieCountThree(string(file))
	fmt.Printf("Day 1 Part 2: %d\n", value2)
}

func day2() {
	file, err := os.ReadFile("input/day02.txt")

	if err != nil {
		panic("Could not read input file " + err.Error())
	}

	value := RpsScore(string(file))
	fmt.Printf("Day 2 Part 1: %d\n", value)

	value2 := RpsPlays(string(file))
	fmt.Printf("Day 2 Part 2: %d\n", value2)
}
