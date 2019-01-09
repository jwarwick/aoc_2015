package main

import (
	"testing"
)

func TestHousesDeliveredTo(t *testing.T) {
	tables := []struct {
		input  string
		result int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, table := range tables {
		result := housesDeliveredTo(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}
