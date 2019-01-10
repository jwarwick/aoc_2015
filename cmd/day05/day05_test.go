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
		result := part1Nice(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %t, got: %t", table.input, table.result, result)
		}
	}
}

func TestPart2Nice(t *testing.T) {
	tables := []struct {
		input  string
		result bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"aaa", false},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, table := range tables {
		result := part2Nice(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %t, got: %t", table.input, table.result, result)
		}
	}
}
