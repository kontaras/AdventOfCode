package main

type Day04 struct{}

var _ Day = Day04{}

// InputPath implements Day.
func (d Day04) InputPath() string {
	return "input/day04.txt"
}

// Part1 implements Day.
func (d Day04) Part1(input string) (int, error) {
	lines := SplitInput(input)

	numLines := len(lines)
	if numLines == 0 {
		return 0, nil
	}
	lineLength := len(lines[0])

	inRange := func(i, j int) bool {
		return i >= 0 && i < numLines && j >= 0 && j < lineLength
	}

	isXmas := func(i0, j0, i1, j1, i2, j2 int) bool {
		return inRange(i0, j0) && lines[i0][j0] == 'M' &&
			inRange(i1, j1) && lines[i1][j1] == 'A' &&
			inRange(i2, j2) && lines[i2][j2] == 'S'
	}

	count := 0

	for i := range numLines {
		for j := range lineLength {
			if lines[i][j] == 'X' {
				if isXmas(i-1, j-1, i-2, j-2, i-3, j-3) {
					count++
				}
				if isXmas(i-1, j, i-2, j, i-3, j) {
					count++
				}

				if isXmas(i-1, j+1, i-2, j+2, i-3, j+3) {
					count++
				}
				if isXmas(i, j-1, i, j-2, i, j-3) {
					count++
				}
				if isXmas(i, j, i, j, i, j) {
					count++
				}

				if isXmas(i, j+1, i, j+2, i, j+3) {
					count++
				}
				if isXmas(i+1, j-1, i+2, j-2, i+3, j-3) {
					count++
				}
				if isXmas(i+1, j, i+2, j, i+3, j) {
					count++
				}

				if isXmas(i+1, j+1, i+2, j+2, i+3, j+3) {
					count++
				}
			}
		}
	}

	return count, nil
}

// Part2 implements Day.
func (d Day04) Part2(input string) (int, error) {
	lines := SplitInput(input)

	numLines := len(lines)
	if numLines == 0 {
		return 0, nil
	}
	lineLength := len(lines[0])

	inRange := func(i, j int) bool {
		return i >= 0 && i < numLines && j >= 0 && j < lineLength
	}

	isXmas := func(i, j int) bool {
		countM := 0
		countS := 0
		for _, coordinate := range [][]int{
			{i - 1, j - 1},
			{i - 1, j + 1},
			{i + 1, j - 1},
			{i + 1, j + 1},
		} {
			x := coordinate[0]
			y := coordinate[1]
			for !inRange(x, y) {
				return false
			}
			if lines[x][y] == 'M' {
				countM++
			}
			if lines[x][y] == 'S' {
				countS++
			}
		}
		return countM == 2 && countS == 2 &&
			lines[i-1][j-1] != lines[i+1][j+1] &&
			lines[i+1][j-1] != lines[i-1][j+1]
	}

	count := 0

	for i := range numLines {
		for j := range lineLength {
			if lines[i][j] == 'A' {
				if isXmas(i, j) {
					count++
				}
			}
		}
	}

	return count, nil
}
