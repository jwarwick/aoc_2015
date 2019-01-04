package main

import (
	"github.com/jwarwick/aoc_2015/internal/day02"
	"testing"
)

func TestPaperSize(t *testing.T) {
	tables := []struct {
		input  string
		result int
	}{
		{"2x3x4", 58},
		{"1x1x10", 43},
	}

	for _, table := range tables {
		result := day02.PaperNeeded(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}

func TestRibbonSize(t *testing.T) {
	tables := []struct {
		input  string
		result int
	}{
		{"2x3x4", 34},
		{"1x1x10", 14},
	}

	for _, table := range tables {
		result := day02.RibbonNeeded(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}
