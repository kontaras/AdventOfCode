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
	fmt.Println("Day 1 Part 1:", Day01Part1(string(file)))
}
