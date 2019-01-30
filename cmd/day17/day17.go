package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	part1 := exactCombinations(string(input), 150)
	fmt.Println("Part 1: ", part1)
}

func exactCombinations(input string, amount int) int {
	lines := splitInput(input)
	cnt := len(lines)
	sizes := make([]int, cnt)
	for idx, l := range lines {
		sizes[idx] = parseInt(l)
	}

	sum := 0
	for i := 0.0; i < math.Pow(2, float64(cnt)); i++ {
		bits := int(i)
		curr := 0
		for _, s := range sizes {
			if 1 == (bits & 0x01) {
				curr += s
			}
			bits = bits >> 1
		}
		if curr == amount {
			sum += 1
		}
	}
	return sum
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return lines
}

func parseInt(input string) int {
	i, err := strconv.Atoi(input)
	if nil != err {
		panic("Not a number when parsing")
	}
	return i
}
