package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	start := 20151125
	tables := []struct {
		start  int
		row    int
		col    int
		result int
	}{
		{start, 1, 1, 20151125},
		{start, 2, 1, 31916031},
		{start, 5, 4, 6899651},
	}

	for _, table := range tables {
		result := code(table.start, table.row, table.col)
		if result != table.result {
			t.Errorf("Input: (%d, %d), expected: %d, got: %d", table.row, table.col, table.result, result)
		}
	}
}
