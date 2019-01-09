package main

import (
	"testing"
)

func TestMinMD5(t *testing.T) {
	tables := []struct {
		input  string
		result int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, table := range tables {
		result := minMD5(table.input, part1Check)
		if result != table.result {
			t.Errorf("Input: %s, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}
