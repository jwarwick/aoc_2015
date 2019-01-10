package main

import (
	"testing"
)

func TestNice(t *testing.T) {
	tables := []struct {
		input  string
		result bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, table := range tables {
		result := nice(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %t, got: %t", table.input, table.result, result)
		}
	}
}
