package main

import (
	"fmt"
	"os"
)

func main() {
	day1()
}

func day1() {
	file, _ := os.ReadFile("input/day01.txt")

	value := CalorieCount(string(file))
	fmt.Printf("Day 1 Part 1: %d\n", value)

	value2 := CalorieCountThree(string(file))
	fmt.Printf("Day 1 Part 2: %d\n", value2)
}
