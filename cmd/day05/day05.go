package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	part1 := countNice(string(input))
	fmt.Println("Part 1: ", part1)
}

func countNice(input string) int {
	total := 0
	lines := splitInput(input)
	for _, l := range lines {
		if nice(l) {
			total++
		}
	}
	return total
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	return lines
}

func nice(input string) bool {
	return hasVowels(input) && hasRepeats(input) && !hasNaughtyLetters(input)
}

func hasVowels(input string) bool {
	total := 0
	for _, c := range input {
		if isVowel(c) {
			total++
		}

		if total >= 3 {
			return true
		}
	}
	return false
}

func isVowel(c rune) bool {
	return c == 'a' ||
		c == 'e' ||
		c == 'i' ||
		c == 'o' ||
		c == 'u'
}

func hasRepeats(input string) bool {
	var last rune
	for _, c := range input {
		if c == last {
			return true
		}
		last = c
	}
	return false
}

func hasNaughtyLetters(input string) bool {
	return strings.Contains(input, "ab") ||
		strings.Contains(input, "cd") ||
		strings.Contains(input, "pq") ||
		strings.Contains(input, "xy")
}
