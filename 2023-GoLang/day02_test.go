package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/kontaras/AdventOfCode/2023-GoLang"
)

var _ = Describe("Day02", func() {
	Context("Part 1", func() {
		It("should parse a cube set", func() {
			Expect(ParseCubePull("3 blue, 4 red")).To(Equal(CubePull{Red: 4, Blue: 3}))
			Expect(ParseCubePull("1 red, 2 green, 6 blue")).To(Equal(CubePull{Red: 1, Green: 2, Blue: 6}))
			Expect(ParseCubePull("2 green")).To(Equal(CubePull{Green: 2}))
		})

		It("should parse a game", func() {
			Expect(ParseGame("3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")).To(Equal([]CubePull{{Red: 4, Blue: 3}, {Red: 1, Green: 2, Blue: 6}, {Green: 2}}))
		})

		It("should find valid games", func() {
			input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
			Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
			Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
			Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
			Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

			Expect(Day02Part1(input)).To(Equal(8))
		})
	})

	Context("Part 2", func() {
		It("should find min cubes", func() {
			Expect(MinCubes(ParseGame("3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))).
				To(Equal(CubePull{Red: 4, Green: 2, Blue: 6}))
			Expect(MinCubes(ParseGame("1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"))).
				To(Equal(CubePull{Red: 1, Green: 3, Blue: 4}))
			Expect(MinCubes(ParseGame("8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"))).
				To(Equal(CubePull{Red: 20, Green: 13, Blue: 6}))
			Expect(MinCubes(ParseGame("1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"))).
				To(Equal(CubePull{Red: 14, Green: 3, Blue: 15}))
			Expect(MinCubes(ParseGame("6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"))).
				To(Equal(CubePull{Red: 6, Green: 3, Blue: 2}))

		})

		It("should calculate pull power", func() {
			Expect(PullPower(CubePull{Red: 4, Green: 2, Blue: 6})).To(Equal(48))
			Expect(PullPower(CubePull{Red: 1, Green: 3, Blue: 4})).To(Equal(12))
			Expect(PullPower(CubePull{Red: 20, Green: 13, Blue: 6})).To(Equal(1560))
			Expect(PullPower(CubePull{Red: 14, Green: 3, Blue: 15})).To(Equal(630))
			Expect(PullPower(CubePull{Red: 6, Green: 3, Blue: 2})).To(Equal(36))
		})

		It("should find valid games", func() {
			input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
			Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
			Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
			Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
			Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

			Expect(Day02Part2(input)).To(Equal(2286))
		})
	})
})
