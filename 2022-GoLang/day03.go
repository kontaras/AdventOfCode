package main

import (
	"strings"
)

func RucksackPriorityScore(manifest string) int {
	lines := strings.Split(manifest, "\n")
	var score int
	for _, sack := range lines {
		item_count := len(sack)
		if item_count > 0 {
			side1 := sack[:item_count/2]
			side2 := sack[item_count/2:]
			var items = make(map[byte]byte)

			for i := 0; i < len(side1); i++ {
				items[side1[i]] = 0
			}

			for i := 0; i < len(side2); i++ {
				item := side2[i]
				_, ok := items[item]
				if ok {
					if item > 'Z' {
						score += (int)(item - 'a' + 1)
					} else {
						score += (int)(item - 'A' + 27)
					}

					break
				}
			}
		}
	}

	return score
}
