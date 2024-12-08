package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day01", func() {
	underTest := Day01{}
	input := `
		3   4
		4   3
		2   5
		1   3
		3   9
		3   3
	`

	Context("Part 1", func() {

		It("should calculate differences", func() {
			left, right, err := underTest.SplitLists(SplitInput(input))
			Expect(err).To(BeNil())
			Expect(left).To(Equal([]int{3, 4, 2, 1, 3, 3}))
			Expect(right).To(Equal([]int{4, 3, 5, 3, 9, 3}))

			Expect(underTest.Differences(right, left)).To(Equal([]int{2, 1, 0, 1, 2, 5}))
		})

		It("should solve the sample", func() {
			Expect(underTest.Part1(input)).To(Equal(11))
		})
	})
	Context("Part 2", func() {
		It("should calculate similarity score", func() {
			left, right, err := underTest.SplitLists(SplitInput(input))
			Expect(err).To(BeNil())

			Expect(underTest.similarityScore(left, right)).To(Equal([]int{9, 4, 0, 0, 9, 9}))
		})

		It("should solve the sample", func() {
			Expect(underTest.Part2(input)).To(Equal(31))
		})
	})
})
