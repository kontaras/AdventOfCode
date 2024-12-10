package main

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "2024 GoLang Suite")
}

var _ = Describe("All days", func() {
	DescribeTable("Day", func(day Day, answer1 int, answer2 int) {
		file, err := os.ReadFile(day.InputPath())
		Expect(err).NotTo(HaveOccurred())
		input := string(file)

		if answer1 != -1 {
			Expect(day.Part1(input)).To(Equal(answer1))
		}

		if answer2 != -1 {
			Expect(day.Part2(input)).To(Equal(answer2))
		}
	},
		Entry("01", Day01{}, 765748, 27732508),
		Entry("02", Day02{}, 236, 308),
		Entry("03", Day03{}, 166357705, 88811886),
		Entry("04", Day04{}, 2483, 1925),
	)
})
