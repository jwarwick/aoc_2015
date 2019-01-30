package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	clues := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1}

	aunts := buildMap(string(input))

	part1 := findAunt(clues, aunts, part1Check)
	fmt.Println("Part 1: ", part1)

	part2 := findAunt(clues, aunts, part2Check)
	fmt.Println("Part 2: ", part2)
}

type check func(string, int, int) bool

func part1Check(_k string, v int, clue_val int) bool {
	return v == clue_val
}

func part2Check(k string, v int, clue_val int) bool {
	valid := true
	switch k {
	case "cats", "trees":
		valid = (v > clue_val)
	case "pomeranians", "goldfish":
		valid = (v < clue_val)
	default:
		valid = (v == clue_val)
	}
	return valid
}

func findAunt(clues map[string]int, aunts []map[string]int, fn check) int {
	for idx, m := range aunts {
		valid := true
		for k, v := range m {
			clue_val, present := clues[k]
			if !present {
				continue
			}
			if !fn(k, v, clue_val) {
				valid = false
				break
			}
		}
		if valid {
			return idx + 1
		}
	}
	return -1
}

func buildMap(input string) []map[string]int {
	lines := splitInput(input)
	aunts := make([]map[string]int, 500)
	for idx, val := range lines {
		m := parseLine(val)
		aunts[idx] = m
	}
	return aunts
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return lines
}

func parseLine(input string) map[string]int {
	values := make(map[string]int)
	words := strings.Split(input, " ")
	for i := 2; i <= 6; i += 2 {
		k := strings.Trim(words[i], ":")
		v := parseInt(strings.Trim(words[i+1], ","))
		values[k] = v
	}
	return values
}

func parseInt(input string) int {
	i, err := strconv.Atoi(input)
	if nil != err {
		panic("Not a number when parsing")
	}
	return i
}
