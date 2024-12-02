package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

var days = map[int]Day{
	1: Day01{},
}

func main() {
	args := os.Args
	if len(args) != 2 || args[1] == "-h" {
		fmt.Printf("Usage: %s DAY_NUM\n", args[0])
		fmt.Println("Available days are:", slices.Sorted(maps.Keys(days)))
		return
	}

	dayNum, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Day number should be numeric, got", args[1])
		return
	}

	day, ok := days[dayNum]
	if !ok {
		fmt.Println("Invalid day", dayNum)
		fmt.Println("Available days are:", slices.Sorted(maps.Keys(days)))
		return
	}

	RunDay(dayNum, day)
}

type Day interface {
	InputPath() string
	Part1(string) (int, error)
	Part2(string) (int, error)
}

func RunDay(tag int, day Day) {
	file, err := os.ReadFile(day.InputPath())
	if err != nil {
		fmt.Printf("Error running Day %d: %s", tag, err)
		return
	}
	input := string(file)

	fmt.Println("Day", tag)

	answer, err := day.Part1(input)
	if err != nil {
		fmt.Println("\tError running part 1", err)
	} else {
		fmt.Println("\tPart 1", answer)
	}

	answer, err = day.Part2(input)
	if err != nil {
		fmt.Println("\tError running part 2", err)
	} else {
		fmt.Println("\tPart 2", answer)
	}
}

func SplitInput(in string) []string {
	lines := []string{}
	for _, line := range strings.Split(in, "\n") {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	return lines
}
