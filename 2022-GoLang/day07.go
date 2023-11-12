package main

import (
	"strconv"
	"strings"
)

var allDirs = []Dir{}

func DirsToDelete(term string) int {
	ParseCmds(term)

	sum := 0

	for _, dir := range allDirs {
		size := dir.Size()
		if size <= 100000 {
			sum += size
		}
	}

	return sum
}

func FindOptimalDelete(term string) int {
	maxUsed := 70000000 - 30000000
	root := ParseCmds(term)
	neededFree := root.Size() - maxUsed
	minDeleted := 70000000
	for _, dir := range allDirs {
		size := dir.Size()
		if size >= neededFree && size < minDeleted {
			minDeleted = size
		}
	}

	return minDeleted
}

func ParseCmds(term string) Dir {
	lines := strings.Split(term, "\n")
	root := *NewDir()
	allDirs = []Dir{root}
	cwd := []Dir{root}
	for _, line := range lines {
		words := strings.Fields(line)
		if len(words) > 0 {
			if (words[0] == "$") {
				cmd := words[1]
				if cmd == "ls" {
					// Nothing
				} else if cmd == "cd" {
					dirName := words[2]
					if dirName == "/" {
						cwd = []Dir{root}
					} else if dirName == ".." {
						cwd = cwd[:len(cwd) - 1]
					} else {
						cwd = append(cwd, cwd[len(cwd) - 1].AddDir(dirName))
					}
				} else {
					panic("Unknown command "+ cmd)
				}
			} else {
				if words[0] == "dir" {
					cwd[len(cwd) - 1].AddDir(words[1])
				} else {
					size, err := strconv.Atoi(words[0])
					if err != nil {
						panic(err)
					}
					cwd[len(cwd) - 1].AddFile(words[1], size)
				}
			}
		}
	}

	return root
}

type File = int
type Dir struct {
	Files map[string]File
	SubDirs map[string]Dir
	size int
}

func (d Dir) AddFile(name string, size int) {
	d.Files[name] = size
}

func (d Dir) AddDir(name string) Dir {
	sub, ok := d.SubDirs[name]
	if ok {
		return sub
	} else {
		d.SubDirs[name] = *NewDir()
		allDirs = append(allDirs, d.SubDirs[name])
		return d.SubDirs[name]
	}
}

func (d Dir) Size() int {
	if d.size >= 0 {
		return d.size
	} else {
		total := 0
		for _, file := range d.Files {
			total += file
		}

		for _, sub := range d.SubDirs {
			total += sub.Size()
		}

		d.size = total

		return total
	}
}

func NewDir() *Dir {
	return &Dir{make(map[string]File),make(map[string]Dir), -1}
}