package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	part1 := memoryDiff(string(input))
	fmt.Println("Part 1: ", part1)
}

func memoryDiff(input string) int {
	literals := 0
	memory := 0
	lines := splitInput(input)

	var double_slash = regexp.MustCompile(`\\\\`)
	var quote = regexp.MustCompile(`\\"`)
	var hex = regexp.MustCompile(`\\x..`)
	for _, l := range lines {
		literals += len(l)
		v1 := double_slash.ReplaceAllString(l, `!`)
		v1 = quote.ReplaceAllString(v1, `@`)
		v1 = hex.ReplaceAllString(v1, `$`)
		memory += len(v1) - 2
	}
	return literals - memory
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return lines
}
