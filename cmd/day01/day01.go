package main

import (
	"fmt"
	"github.com/jwarwick/aoc_2015/internal/day01"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	part1 := day01.FloorCount(string(input))
	fmt.Println("Part 1: ", part1)

	part2 := day01.Basement(string(input))
	fmt.Println("Part 2: ", part2)
}
