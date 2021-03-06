package main

import (
	"github.com/jwarwick/aoc_2015/internal/day01"
	"testing"
)

func TestFloorCount(t *testing.T) {
	tables := []struct {
		input  string
		result int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, table := range tables {
		result := day01.FloorCount(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}

func TestBasement(t *testing.T) {
	tables := []struct {
		input  string
		result int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, table := range tables {
		result := day01.Basement(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}
