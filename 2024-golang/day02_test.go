package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day02", func() {
	underTest := Day02{}
	input := `
		7 6 4 2 1
		1 2 7 8 9
		9 7 6 2 1
		1 3 2 4 5
		8 6 4 4 1
		1 3 6 7 9
	`

	// 52 51 50 49 46 39 39 false

	Context("Part 1", func() {
		It("should identify safe levels", func() {
			Expect(underTest.isSafe("7 6 4 2 1")).To(BeTrue())
			Expect(underTest.isSafe("1 2 4 6 7")).To(BeTrue())
			Expect(underTest.isSafe("1 2 7 8 9")).To(BeFalse())
			Expect(underTest.isSafe("9 7 6 2 1")).To(BeFalse())
			Expect(underTest.isSafe("1 3 2 4 5")).To(BeFalse())
			Expect(underTest.isSafe("8 6 4 4 1")).To(BeFalse())
			Expect(underTest.isSafe("1 3 6 7 9")).To(BeTrue())
			Expect(underTest.isSafe("9 7 6 3 1")).To(BeTrue())

			Expect(underTest.isSafe("26 27 28 31 33 34 37 37")).To(BeFalse())
			Expect(underTest.isSafe("92 94 97 98 97")).To(BeFalse())
			Expect(underTest.isSafe("56 59 60 61 62 65 69")).To(BeFalse())
			Expect(underTest.isSafe("8 11 13 16 19 20 23 24")).To(BeTrue())
		})

		It("should solve the sample", func() {
			Expect(underTest.Part1(input)).To(Equal(2))
		})
	})
})
