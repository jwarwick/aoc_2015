package main

import (
	"testing"
)

func TestRace(t *testing.T) {
	input := `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

	result := distance(input, 1000)
	if result != 1120 {
		t.Errorf("Expected: 1120, got: %d", result)
	}
}
