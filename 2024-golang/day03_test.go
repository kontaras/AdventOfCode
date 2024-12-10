package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day03", func() {
	underTest := Day03{}

	Context("Part 1", func() {
		input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

		It("should parse muls", func() {
			Expect(underTest.findMuls(input)).To(Equal([]mul{
				{2, 4},
				{5, 5},
				{11, 8},
				{8, 5},
			}))
		})

		It("should solve the sample", func() {
			Expect(underTest.Part1(input)).To(Equal(161))
		})
	})

	Context("Part 2", func() {
		input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

		It("should parse muls", func() {
			Expect(underTest.findEnabledMuls(input)).To(Equal([]mul{
				{2, 4},
				{8, 5},
			}))
		})

		It("should solve the sample", func() {
			Expect(underTest.Part2(input)).To(Equal(48))
		})
	})
})
