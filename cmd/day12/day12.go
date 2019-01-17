package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	part1 := stringSum(string(input))
	fmt.Println("Part 1: ", part1)
}

func stringSum(input string) int {
	var numbers = regexp.MustCompile(`[-]?\d{1,}`)
	matches := numbers.FindAllString(input, -1)
	total := 0
	for _, s := range matches {
		val, err := strconv.Atoi(s)
		if nil != err {
			fmt.Println("Invalid: ", val)
			panic("Not a number")
		}
		total += val
	}
	return total
}
