package main

import (
	"fmt"
	"github.com/jwarwick/aoc_2015/internal/day02"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	part1 := day02.WrappingPaperSum(string(input))
	fmt.Println("Part 1: ", part1)

	part2 := day02.RibbonSum(string(input))
	fmt.Println("Part 2: ", part2)
}
