package main

import (
	"testing"
)

func TestDay01(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`
	actual := CalorieCount(input)

	if actual != 24000 {
		t.Error(actual)
	}

	actual2 := CalorieCountThree(input)

	if actual2 != 45000 {
		t.Error(actual2)
	}
}

func TestDay02(t *testing.T) {
	input := `A Y
B X
C Z`
	actual := RpsScore(input)
	if actual != 15 {
		t.Error(actual)
	}

	actual2 := RpsPlays(input)
	if actual2 != 12 {
		t.Error(actual2)
	}
}

func TestDay03(t *testing.T) {
	var input = `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`
	var actual = RucksackPriorityScore(input)

	if actual != 157 {
		t.Error(actual)
	}

	input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
`

	actual = GroupBadge(input)

	if actual != 18 {
		t.Error(actual)
	}

	input = `wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	ttgJtRGJQctTZtZT
	CrZsJsPPZsGzwwsLwLmpwMDw
`

	actual = GroupBadge(input)

	if actual != 52 {
		t.Error(actual)
	}

	input = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

	actual = GroupBadge(input)

	if actual != 70 {
		t.Error(actual)
	}
}

func TestDay04(t *testing.T) {
	input := `
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`
	var actual = OverlappingAssignments(input)
	if actual != 2 {
		t.Error(actual)
	}

	// actual = RpsPlays(input)
	// if actual != 12 {
	// 	t.Error(actual)
	// }
}
