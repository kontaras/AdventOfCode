package main

import (
	"strings"
)

func RucksackPriorityScore(manifest string) int {
	lines := strings.Split(manifest, "\n")
	var score int
	for _, sack := range lines {
		itemCount := len(sack)
		if itemCount > 0 {
			side1 := sack[:itemCount/2]
			side2 := sack[itemCount/2:]
			var items = make(map[byte]byte)

			for i := 0; i < len(side1); i++ {
				items[side1[i]] = 0
			}

			for i := 0; i < len(side2); i++ {
				item := side2[i]
				_, ok := items[item]
				if ok {
					score += itemPriority(item)

					break
				}
			}
		}
	}

	return score
}

func GroupBadge(elves string) int {
	lines := strings.Split(elves, "\n")
	var score int

	const groupSize = 3

	for i := 0; i+groupSize < len(lines); i += groupSize {
		group := lines[i : i+groupSize]

		var items = make(map[byte]byte)

		for i := 0; i < len(group[0]); i++ {
			items[group[0][i]] = 0
		}

		for _, member := range group[1:] {
			itemsNew := make(map[byte]byte)
			for i := 0; i < len(member); i++ {
				_, ok := items[member[i]]
				if ok {
					itemsNew[member[i]] = 0
				}
			}

			items = itemsNew
		}

		for item := range items {
			score += itemPriority(item)
		}
	}

	return score
}

func itemPriority(item byte) int {
	if item > 'Z' {
		return (int)(item - 'a' + 1)
	} else {
		return (int)(item - 'A' + 27)
	}
}
