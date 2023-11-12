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
	var actual = SubsetAssignments(input)
	if actual != 2 {
		t.Error(actual)
	}

	actual = OverlappingAssignments(input)
	if actual != 4 {
		t.Error(actual)
	}
}

func TestDay05(t *testing.T) {
	input := `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`
	var actual = CrateStack(input)
	if actual != "CMZ" {
		t.Error(actual)
	}

	actual = CrateStack2(input)
	if actual != "MCD" {
		t.Error(actual)
	}
}

func TestDay06(t *testing.T) {
	var actual = FindPacketStart("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	if actual != 7 {
		t.Error(actual)
	}

	actual = FindPacketStart("bvwbjplbgvbhsrlpgdmjqwftvncz")
	if actual != 5 {
		t.Error(actual)
	}

	actual = FindPacketStart("nppdvjthqldpwncqszvftbrmjlhg")
	if actual != 6 {
		t.Error(actual)
	}

	actual = FindPacketStart("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	if actual != 10 {
		t.Error(actual)
	}

	actual = FindPacketStart("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	if actual != 11 {
		t.Error(actual)
	}

	actual = FindMessageStart("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	if actual != 19 {
		t.Error(actual)
	}

	actual = FindMessageStart("bvwbjplbgvbhsrlpgdmjqwftvncz")
	if actual != 23 {
		t.Error(actual)
	}

	actual = FindMessageStart("nppdvjthqldpwncqszvftbrmjlhg")
	if actual != 23 {
		t.Error(actual)
	}

	actual = FindMessageStart("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	if actual != 29 {
		t.Error(actual)
	}

	actual = FindMessageStart("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	if actual != 26 {
		t.Error(actual)
	}
}

func TestDay07(t *testing.T) {
	input := `$ cd /
	$ ls
	dir a
	14848514 b.txt
	8504156 c.dat
	dir d
	$ cd a
	$ ls
	dir e
	29116 f
	2557 g
	62596 h.lst
	$ cd e
	$ ls
	584 i
	$ cd ..
	$ cd ..
	$ cd d
	$ ls
	4060174 j
	8033020 d.log
	5626152 d.ext
	7214296 k
	`

	actual := DirsToDelete(input)
	if actual != 95437 {
		t.Error(actual)
	}

	actual = FindOptimalDelete(input)
	if actual != 24933642 {
		t.Error(actual)
	}
}
