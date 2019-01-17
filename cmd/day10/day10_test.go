package main

import (
	"testing"
)

func TestLookAndSay(t *testing.T) {
	tables := []struct {
		input  string
		result string
	}{
		{"1", "11"},
		{"11", "21"},
		{"1211", "111221"},
		{"111221", "312211"},
	}

	for _, table := range tables {
		result := lookAndSay(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %s, got: %s", table.input, table.result, result)
		}
	}

	result := lookAndSayMulti("1", 5)
	if "312211" != result {
		t.Errorf("Multi input: 1, expected: 312211, got: %s", result)
	}
}
