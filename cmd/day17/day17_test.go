package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `20
15
10
5
5`

	count, minWays := exactCombinations(input, 25)
	if count != 4 {
		t.Errorf("TotalWays Expected: 4, got: %d", count)
	}
	if minWays != 3 {
		t.Errorf("MinWays Expected: 3, got: %d", minWays)
	}
}
