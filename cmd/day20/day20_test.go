package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tables := []struct {
		input  int
		result int
	}{
		{10, 1},
		{40, 3},
		{80, 6},
		{61, 4},
		{119, 6},
	}

	for _, table := range tables {
		result := findHouse(table.input)
		if result != table.result {
			t.Errorf("Input: %d, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}
