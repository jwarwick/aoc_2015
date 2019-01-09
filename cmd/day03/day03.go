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
	part1 := housesDeliveredTo(string(input), 1)
	fmt.Println("Part 1: ", part1)

	part2 := housesDeliveredTo(string(input), 2)
	fmt.Println("Part 2: ", part2)
}

func housesDeliveredTo(input string, num int) int {
	houses := make(map[complex128]int, len(input))
	locs := make([]complex128, num)
	for i := 0; i < num; i++ {
		locs[i] = 0 + 0i
		houses[locs[i]] += 1
	}
	curr_loc := 0

	for _, c := range input {
		idx := curr_loc % len(locs)
		curr := locs[idx]
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
		locs[idx] = curr
		curr_loc += 1
	}
	return len(houses)
}
