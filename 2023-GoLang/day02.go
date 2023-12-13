package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day02Part1(input string) int {
	sum := 0

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line := strings.Trim(line, " \t")
		if len(line) < 1 {
			continue
		}

		gameName, gameConnent, cut := strings.Cut(line, ": ")
		if !cut {
			panic("Could not parse game " + line)
		}
		var gameId int
		n, err := fmt.Sscanf(gameName, "Game %d", &gameId)
		if err != nil || n != 1 {
			panic("Could not parse game ID:" + gameName + err.Error())
		}

		pulls := ParseGame(gameConnent)
		if ValidateGame(pulls) {
			sum += gameId
		}
	}

	return sum
}

type CubePull struct {
	Red   int
	Green int
	Blue  int
}

func ParseCubePull(pullStr string) CubePull {
	pull := CubePull{}
	cubeSets := strings.Split(pullStr, ", ")
	for _, cubeSet := range cubeSets {
		num, color, cut := strings.Cut(cubeSet, " ")
		if !cut {
			panic("Could not parse cubeSet " + cubeSet)
		}

		quantity, err := strconv.Atoi(num)

		if err != nil {
			panic("Could not parse number " + err.Error())
		}

		switch color {
		case "red":
			pull.Red = quantity
		case "green":
			pull.Green = quantity
		case "blue":
			pull.Blue = quantity
		default:
			panic("Unknown cube color " + color)
		}
	}
	return pull
}

func ParseGame(game string) []CubePull {
	pullStrs := strings.Split(game, "; ")
	pulls := make([]CubePull, len(pullStrs))
	for i, pull := range pullStrs {
		pulls[i] = ParseCubePull(pull)
	}

	return pulls
}

func ValidateGame(pulls []CubePull) bool {
	for _, pull := range pulls {
		if pull.Red > 12 || pull.Green > 13 || pull.Blue > 14 {
			return false
		}
	}
	return true
}
