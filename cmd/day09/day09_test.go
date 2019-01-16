package main

import (
	"testing"
)

func TestMinMaxPath(t *testing.T) {
	input := `London to Dublin = 464
  London to Belfast = 518
  Dublin to Belfast = 141`

	shortest, longest := minMaxPath(input)
	expected_shortest := 605
	expected_longest := 982
	if shortest != expected_shortest {
		t.Errorf("Shortest Expected: %d, got: %d", expected_shortest, shortest)
	}
	if longest != expected_longest {
		t.Errorf("Longest Expected: %d, got: %d", expected_longest, longest)
	}
}
