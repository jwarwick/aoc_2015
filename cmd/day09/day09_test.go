package main

import (
	"testing"
)

func TestShortestPath(t *testing.T) {
	input := `London to Dublin = 464
  London to Belfast = 518
  Dublin to Belfast = 141`

	result := shortestPath(input)
	expected := 605
	if result != expected {
		t.Errorf("Expected: %d, got: %d", expected, result)
	}
}
