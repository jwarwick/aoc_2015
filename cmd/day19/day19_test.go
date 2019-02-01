package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	rules := []rule{rule{"H", "HO"}, rule{"H", "OH"}, rule{"O", "HH"}}

	tables := []struct {
		input  string
		result int
	}{
		{"HOH", 4},
		{"HOHOHO", 7},
	}

	for _, table := range tables {
		result := molecules(table.input, rules)
		if result != table.result {
			t.Errorf("Input: %s, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}

func TestPart2(t *testing.T) {
	rules := []rule{rule{"e", "H"}, rule{"e", "O"}, rule{"H", "HO"}, rule{"H", "OH"}, rule{"O", "HH"}}

	tables := []struct {
		input  string
		result int
	}{
		{"HOH", 3},
		{"HOHOHO", 6},
	}

	for _, table := range tables {
		result := build(table.input, rules)
		if result != table.result {
			t.Errorf("Input: %s, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}
