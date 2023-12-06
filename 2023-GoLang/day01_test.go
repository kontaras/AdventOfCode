package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	main "github.com/kontaras/AdventOfCode/2023-GoLang"
)

var _ = Describe("Day01", func() {
	Context("Part 1", func() {
		It("should calibrate a line", func() {
			Expect(main.CalibrateLine("1abc2")).To(Equal(12))
			Expect(main.CalibrateLine("pqr3stu8vwx")).To(Equal(38))
			Expect(main.CalibrateLine("a1b2c3d4e5f")).To(Equal(15))
			Expect(main.CalibrateLine("treb7uchet")).To(Equal(77))

			Expect(main.CalibrateLine("")).To(Equal(0))
		})

		It("should handle example input", func() {
			var input = `1abc2
			pqr3stu8vwx
			a1b2c3d4e5f
			treb7uchet`
			Expect(main.Day01Part1(input)).To(Equal(142))
		})
	})

	Context("Part 2", func() {
		It("should find numbers at start of string", func() {
			Expect(main.ParseNumber("foo")).To(Equal(-1))
			Expect(main.ParseNumber("0")).To(Equal(0))
			Expect(main.ParseNumber("1cat")).To(Equal(1))
			Expect(main.ParseNumber("2three")).To(Equal(2))
			Expect(main.ParseNumber("31")).To(Equal(3))
			Expect(main.ParseNumber("4is a thing")).To(Equal(4))
			Expect(main.ParseNumber("5")).To(Equal(5))
			Expect(main.ParseNumber("6")).To(Equal(6))
			Expect(main.ParseNumber("7")).To(Equal(7))
			Expect(main.ParseNumber("8")).To(Equal(8))
			Expect(main.ParseNumber("9")).To(Equal(9))
			Expect(main.ParseNumber(" 5")).To(Equal(-1))
			Expect(main.ParseNumber("zero")).To(Equal(0))
			Expect(main.ParseNumber("ones")).To(Equal(1))
			Expect(main.ParseNumber("twothree")).To(Equal(2))
			Expect(main.ParseNumber("three11")).To(Equal(3))
			Expect(main.ParseNumber("fourteen")).To(Equal(4))
			Expect(main.ParseNumber("fiver")).To(Equal(5))
			Expect(main.ParseNumber("six ")).To(Equal(6))
			Expect(main.ParseNumber("seven\n\t")).To(Equal(7))
			Expect(main.ParseNumber("eight")).To(Equal(8))
			Expect(main.ParseNumber("nine")).To(Equal(9))
		})

		It("should calibrate a line with numeric words", func() {
			Expect(main.CalibrateLineWords("1abc2")).To(Equal(12))
			Expect(main.CalibrateLineWords("pqr3stu8vwx")).To(Equal(38))
			Expect(main.CalibrateLineWords("a1b2c3d4e5f")).To(Equal(15))
			Expect(main.CalibrateLineWords("treb7uchet")).To(Equal(77))

			Expect(main.CalibrateLineWords("")).To(Equal(0))

			Expect(main.CalibrateLineWords("two1nine")).To(Equal(29))
			Expect(main.CalibrateLineWords("eightwothree")).To(Equal(83))
			Expect(main.CalibrateLineWords("abcone2threexyz")).To(Equal(13))
			Expect(main.CalibrateLineWords("xtwone3four")).To(Equal(24))
			Expect(main.CalibrateLineWords("4nineeightseven2")).To(Equal(42))
			Expect(main.CalibrateLineWords("zoneight234")).To(Equal(14))
			Expect(main.CalibrateLineWords("7pqrstsixteen")).To(Equal(76))
		})

		It("should handle example input", func() {
			var input = `1abc2
			pqr3stu8vwx
			a1b2c3d4e5f
			treb7uchet`
			Expect(main.Day01Part2(input)).To(Equal(142))

			input = `two1nine
			eightwothree
			abcone2threexyz
			xtwone3four
			4nineeightseven2
			zoneight234
			7pqrstsixteen`
			Expect(main.Day01Part2(input)).To(Equal(281))
		})
	})
})
