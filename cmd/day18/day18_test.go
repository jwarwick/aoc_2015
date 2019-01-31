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

	count := simulate(input, 4, false)
	if count != 4 {
		t.Errorf("Expected: 4, got: %d", count)
	}
}

func TestPart2(t *testing.T) {
	input := `.#.#.#
...##.
#....#
..#...
#.#..#
####..`

	count := simulate(input, 5, true)
	if count != 17 {
		t.Errorf("Expected: 17, got: %d", count)
	}
}
