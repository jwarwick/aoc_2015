package main

import (
	"testing"
)

func TestMemory(t *testing.T) {
	input := `""
  "abc"
  "a\\bc"
  "aaa\"aaa"
  "\x27"`

	result := memoryDiff(input)
	expected := 15
	if result != expected {
		t.Errorf("Expected: %d, got: %d", expected, result)
	}
}
