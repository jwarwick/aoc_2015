package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `inc a
jio a, +2
tpl a
inc a`

	p := createProgram(input)
	p.run()
	if p.registers[0] != 2 {
		t.Errorf("Expected: 2, got: %d", p.registers[0])
	}
}
