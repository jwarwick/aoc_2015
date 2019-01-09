package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	part1 := housesDeliveredTo(string(input))
	fmt.Println("Part 1: ", part1)

	// part2 := day02.RibbonSum(string(input))
	// fmt.Println("Part 2: ", part2)
}

func housesDeliveredTo(input string) int {
	houses := make(map[complex128]int)
	curr := 0 + 0i
	houses[curr] = 1
	for _, c := range input {
		switch c {
		case '^':
			curr += 0 + 1i
		case 'v':
			curr += 0 - 1i
		case '<':
			curr += -1 + 0i
		case '>':
			curr += 1 + 0i
		}
		houses[curr] += 1
	}
	return len(houses)
}
