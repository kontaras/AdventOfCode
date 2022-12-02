package main

import (
	"testing"
)

func TestDay01(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`
	actual := CalorieCount(input)

	if actual != 24000 {
		t.Error(actual)
	}

	actual2 := CalorieCountThree(input)

	if actual2 != 45000 {
		t.Error(actual2)
	}
}
