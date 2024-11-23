package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/kontaras/AdventOfCode/2023-GoLang"
)

var _ = Describe("Day03", func() {
	Context("Part 1", func() {
		input := `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
		`

		It("should find numbers", func() {
			Expect(GetNumbers(input)).To(Equal([]int{467, 35, 633, 617, 592, 755, 664, 598}))
		})

		It("should get right answer", func() {
			Expect(Day03Part1(input)).To(Equal(4361))
		})
	})

})
