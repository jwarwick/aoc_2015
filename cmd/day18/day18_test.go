package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := `.#.#.#
...##.
#....#
..#...
#.#..#
####..`

	count := simulate(input, 4)
	if count != 4 {
		t.Errorf("Expected: 4, got: %d", count)
	}
}
