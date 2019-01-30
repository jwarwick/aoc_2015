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

	count := exactCombinations(input, 25)
	if count != 4 {
		t.Errorf("Expected: 4, got: %d", count)
	}
}

// func TestCombinations(t *testing.T) {
// 	tables := []struct {
// 		input  []int
// 		result [][]int
// 	}{
// 		{[]int, [][]int{[]int}}
// 		// {[1], [[], [1]]},
// 		// {[1, 2], [[], [1], [2], [1,2]]},
// 	}
//
// 	for _, table := range tables {
// 		result := combinations(table.input)
// 		if result != table.result {
// 			t.Errorf("Input: %d, expected: %s, got: %s", table.input, table.result, result)
// 		}
// 	}
// }
