package main

import (
	"testing"
)

func TestStringSum(t *testing.T) {
	tables := []struct {
		input  string
		result int
	}{
		{`[1,2,3]`, 6},
		{`{"a":2,"b":4}`, 6},
		{`[[[3]]]`, 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`[-55, {"a":56}]`, 1},
		{`[]`, 0},
		{`{}`, 0},
	}

	for _, table := range tables {
		result := stringSum(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}

func TestNotRed(t *testing.T) {
	tables := []struct {
		input  string
		result int
	}{
		{`[1,2,3]`, 6},
		{`{"a":2,"b":4}`, 6},
		{`[[[3]]]`, 3},
		{`{"a":{"b":4},"c":-1}`, 3},
		{`{"a":[-1,1]}`, 0},
		{`[-1,{"a":1}]`, 0},
		{`[-55, {"a":56}]`, 1},
		{`[]`, 0},
		{`{}`, 0},
		{`[1,{"c":"red","b":2},3]`, 4},
		{`{"d":"red","e":[1,2,3,4],"f":5}`, 0},
		{`[1,"red",5]`, 6},
	}

	for _, table := range tables {
		result := notRedStringSum(table.input)
		if result != table.result {
			t.Errorf("Input: %s, expected: %d, got: %d", table.input, table.result, result)
		}
	}
}
