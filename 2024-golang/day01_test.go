package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day01", func() {
	underTest := Day01{}

	Context("Part 1", func() {
		input := `
		3   4
		4   3
		2   5
		1   3
		3   9
		3   3
		`

		It("should calculate differences", func() {
			left, right, err := underTest.SplitLists(SplitInput(input))
			Expect(err).To(BeNil())
			Expect(left).To(Equal([]int{3, 4, 2, 1, 3, 3}))
			Expect(right).To(Equal([]int{4, 3, 5, 3, 9, 3}))

			Expect(underTest.Diffences(right, left)).To(Equal([]int{2, 1, 0, 1, 2, 5}))
		})

		It("should solve the sample", func() {
			Expect(underTest.Part1(input)).To(Equal(11))
		})
	})
})
