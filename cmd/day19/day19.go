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

	rules, start := parseFile(string(input))
	part1 := molecules(start, rules)
	fmt.Println("Part 1: ", part1)
}

type rule struct {
	from string
	to   string
}

func parseFile(input string) (rules []rule, start string) {
	lines := splitInput(input)
	rules = make([]rule, len(lines)-2)
	blankIdx := len(rules)
	startIdx := len(rules) + 1
	for idx, l := range lines {
		switch idx {
		case startIdx:
			start = l
		case blankIdx:
			continue
		default:
			words := strings.Fields(l)
			rules[idx] = rule{words[0], words[2]}
		}
	}
	return
}

func molecules(start string, rules []rule) int {
	set := make(map[string]int)
	for _, r := range rules {
		count := strings.Count(start, r.from)
		for n := 1; n <= count; n++ {
			splits := strings.SplitAfterN(start, r.from, n)
			splits[n-1] = strings.Replace(splits[n-1], r.from, r.to, 1)
			combined := strings.Join(splits, "")
			set[combined]++
		}
	}
	return len(set)
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return lines
}
