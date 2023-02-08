package main

import (
	//"strconv"
	"strings"
)

func DirsToDelete(term string) int {
	lines := strings.Split(term, "\n")
	for _, cmd := range lines {
		if cmd != "" {
			if (cmd[0] == '$') {
				//Command
			} else {
				//output
			}
		}
	}

	return len(lines)
}