package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day04", func() {
	underTest := Day04{}

	Context("Part 1", func() {
		It("should solve the samples", func() {
			input1 := `
				..X...
				.SAMX.
				.A..A.
				XMAS.S
				.X....
			`
			Expect(underTest.Part1(input1)).To(Equal(4))

			input2 := `
				MMMSXXMASM
				MSAMXMSMSA
				AMXSXMAAMM
				MSAMASMSMX
				XMASAMXAMM
				XXAMMXXAMA
				SMSMSASXSS
				SAXAMASAAA
				MAMMMXMMMM
				MXMXAXMASX
			`
			Expect(underTest.Part1(input2)).To(Equal(18))
		})
	})

	Context("Part 2", func() {
		It("should solve the samples", func() {
			input1 := `
					M.S
					.A.
					M.S
				`
			Expect(underTest.Part2(input1)).To(Equal(1))

			input2 := `
					MMMSXXMASM
					MSAMXMSMSA
					AMXSXMAAMM
					MSAMASMSMX
					XMASAMXAMM
					XXAMMXXAMA
					SMSMSASXSS
					SAXAMASAAA
					MAMMMXMMMM
					MXMXAXMASX
				`
			Expect(underTest.Part2(input2)).To(Equal(9))
		})
	})
})
