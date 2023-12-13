package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("input/day01.txt")

	if err != nil {
		panic("Could not read input file: " + err.Error())
	}
	fmt.Println("Day 1:", Day01Part1(string(file)), Day01Part2(string(file)))

	file, err = os.ReadFile("input/day02.txt")

	if err != nil {
		panic("Could not read input file: " + err.Error())
	}
	fmt.Println("Day 2:", Day02Part1(string(file)))
}
