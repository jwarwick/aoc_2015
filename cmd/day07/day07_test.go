package main

import (
	"testing"
)

func TestCircuit(t *testing.T) {
	input := `123 -> x
  456 -> y
  x AND y -> d
  x OR y -> e
  x LSHIFT 2 -> f
  y RSHIFT 2 -> g
  y -> z
  NOT x -> h
  NOT y -> i`

	vars := []struct {
		input  string
		result uint16
	}{
		{"d", 72},
		{"e", 507},
		{"f", 492},
		{"g", 114},
		{"h", 65412},
		{"i", 65079},
		{"x", 123},
		{"y", 456},
		{"z", 456},
	}

	for _, v := range vars {
		result := simulate(input, v.input, make(map[string]uint16))
		if result[v.input] != v.result {
			t.Errorf("Variable: %s, expected: %d, got: %d", v.input, v.result, result[v.input])
		}
	}
}
