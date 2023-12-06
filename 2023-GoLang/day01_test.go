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
})
