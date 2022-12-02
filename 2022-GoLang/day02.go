package main

import (
	"fmt"
	"strings"
)

const win = 6
const draw = 3
const lose = 0

func RpsScore(rounds string) int {
	score := 0
	lines := strings.Split(rounds, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			opponent := parts[0]
			self := parts[1]
			score += moveScore(self)

			score += rpsOutcome(self, opponent)
		}
	}

	return score
}

func RpsPlays(rounds string) int {
	score := 0
	lines := strings.Split(rounds, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			opponent := parts[0]
			outcomeCode := parts[1]

			var outcome int

			switch outcomeCode {
			case "X":
				outcome = lose
			case "Y":
				outcome = draw
			case "Z":
				outcome = win
			}

			score += outcome

			score += moveScore(getMove(opponent, outcome))
		}
	}

	return score
}

func moveScore(play string) int {
	switch play {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}

	panic("Unknown move " + play)
}

func rpsOutcome(self string, opponent string) int {
	switch self {
	case "X":
		switch opponent {
		case "A":
			return draw
		case "B":
			return lose
		case "C":
			return win
		}
	case "Y":
		switch opponent {
		case "A":
			return win
		case "B":
			return draw
		case "C":
			return lose
		}
	case "Z":
		switch opponent {
		case "A":
			return lose
		case "B":
			return win
		case "C":
			return draw
		}
	}

	panic(fmt.Sprintf("Unknown outcome %s %s", self, opponent))
}

func getMove(opponent string, outcome int) string {
	switch opponent {
	case "A":
		switch outcome {
		case win:
			return "Y"
		case draw:
			return "X"
		case lose:
			return "Z"
		}
	case "B":
		switch outcome {
		case win:
			return "Z"
		case draw:
			return "Y"
		case lose:
			return "X"
		}
	case "C":
		switch outcome {
		case win:
			return "X"
		case draw:
			return "Z"
		case lose:
			return "Y"
		}
	}

	panic(fmt.Sprintf("Unknown outcome %s %d", opponent, outcome))
}
