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

		gameName, gameContent, cut := strings.Cut(line, ": ")
		if !cut {
			panic("Could not parse game " + line)
		}
		var gameId int
		n, err := fmt.Sscanf(gameName, "Game %d", &gameId)
		if err != nil || n != 1 {
			panic("Could not parse game ID:" + gameName + err.Error())
		}

		pulls := ParseGame(gameContent)
		if ValidateGame(pulls) {
			sum += gameId
		}
	}

	return sum
}

func Day02Part2(input string) int {
	sum := 0

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line := strings.Trim(line, " \t")
		if len(line) < 1 {
			continue
		}

		gameName, gameContent, cut := strings.Cut(line, ": ")
		if !cut {
			panic("Could not parse game " + line)
		}
		var gameId int
		n, err := fmt.Sscanf(gameName, "Game %d", &gameId)
		if err != nil || n != 1 {
			panic("Could not parse game ID:" + gameName + err.Error())
		}

		pulls := ParseGame(gameContent)
		min := MinCubes(pulls)
		sum += PullPower(min)
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

func MinCubes(pulls []CubePull) CubePull {
	max := CubePull{}
	for _, pull := range pulls {
		if pull.Red > max.Red {
			max.Red = pull.Red
		}

		if pull.Green > max.Green {
			max.Green = pull.Green
		}

		if pull.Blue > max.Blue {
			max.Blue = pull.Blue
		}
	}

	return max
}

func PullPower(pull CubePull) int {
	return pull.Red * pull.Green * pull.Blue
}
