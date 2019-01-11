package main

import (
	"testing"
)

func TestNice(t *testing.T) {
	tables := []struct {
		input  string
		result int
	}{
		{"turn on 0,0 through 999,999", 1000000},
		{"toggle 0,0 through 999,0", 1000},
		{"turn off 499,499 through 500,500", 0}, // turns off the 4 middle lights
	}

	for _, table := range tables {
		result := countLights(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}
