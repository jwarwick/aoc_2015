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

	part1 := countNice(string(input), part1Nice)
	fmt.Println("Part 1: ", part1)

	part2 := countNice(string(input), part2Nice)
	fmt.Println("Part 2: ", part2)
}

type nice_fun func(string) bool

func countNice(input string, f nice_fun) int {
	total := 0
	lines := splitInput(input)
	for _, l := range lines {
		if f(l) {
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

func part1Nice(input string) bool {
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

func part2Nice(input string) bool {
	return repeatedPair(input) && spacedDuplicate(input)
}

func repeatedPair(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		sub := input[i : i+2]
		count := strings.Count(input, sub)
		if count >= 2 {
			return true
		}
	}
	return false
}

func spacedDuplicate(input string) bool {
	var two_back rune
	var last rune
	for _, c := range input {
		if c == two_back {
			return true
		}
		two_back = last
		last = c
	}
	return false
}
