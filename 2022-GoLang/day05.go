package main

import (
	"regexp"
	"strconv"
	"strings"
)

func CrateStack(manifest string) string {
	lines := strings.Split(manifest, "\n")
	var i = 0
	for ; len(lines[i]) == 0; i++ {
	}

	numCrates := len(lines[i])/4 + 1

	stacks := make([]stack, numCrates)

	for ; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], " 1") {
			i += 2
			break
		}

		line := lines[i]
		for j := 0; j < len(line); j += 4 {
			if line[j] == '[' {
				stacks[j/4] = append(stacks[j/4], string(line[j+1]))
			}
		}
	}

	for _, aStack := range stacks {
		for i, j := 0, len(aStack)-1; i < j; i, j = i+1, j-1 {
			aStack[i], aStack[j] = aStack[j], aStack[i]
		}
	}

	r, err := regexp.Compile(`move (\d+) from (\d) to (\d)`)

	if err != nil {
		panic("Could not compile regex: " + err.Error())
	}

	for ; i < len(lines); i++ {
		line := lines[i]
		if match := r.FindStringSubmatch(line); match != nil {
			quantity, _ := strconv.Atoi(match[1])
			var source, _ = strconv.Atoi(match[2])
			var dest, _ = strconv.Atoi(match[3])
			for i := 0; i < quantity; i++ {
				var item string
				stacks[source-1], item = stacks[source-1].Pop()
				stacks[dest-1] = stacks[dest-1].Push(item)
			}
		}
	}

	var tops = ""
	for _, aStack := range stacks {
		tops += aStack[len(aStack)-1]
	}

	return tops
}

func CrateStack2(manifest string) string {
	lines := strings.Split(manifest, "\n")
	var i = 0
	for ; len(lines[i]) == 0; i++ {
	}

	numCrates := len(lines[i])/4 + 1

	stacks := make([]stack, numCrates)

	for ; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], " 1") {
			i += 2
			break
		}

		line := lines[i]
		for j := 0; j < len(line); j += 4 {
			if line[j] == '[' {
				stacks[j/4] = append(stacks[j/4], string(line[j+1]))
			}
		}
	}

	for _, aStack := range stacks {
		for i, j := 0, len(aStack)-1; i < j; i, j = i+1, j-1 {
			aStack[i], aStack[j] = aStack[j], aStack[i]
		}
	}

	r, err := regexp.Compile(`move (\d+) from (\d) to (\d)`)

	if err != nil {
		panic("Could not compile regex: " + err.Error())
	}

	for ; i < len(lines); i++ {
		line := lines[i]
		if match := r.FindStringSubmatch(line); match != nil {
			quantity, _ := strconv.Atoi(match[1])
			var source, _ = strconv.Atoi(match[2])
			var dest, _ = strconv.Atoi(match[3])
			var item []string
			stacks[source-1], item = stacks[source-1].MultiPop(quantity)
			stacks[dest-1] = stacks[dest-1].MultiPush(item)
		}
	}

	var tops = ""
	for _, aStack := range stacks {
		tops += aStack[len(aStack)-1]
	}

	return tops
}

// Stack implementation taken from
type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) MultiPush(v []string) stack {
	return append(s, v...)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) MultiPop(num int) (stack, []string) {
	l := len(s)
	return s[:l-num], s[l-num:]
}
