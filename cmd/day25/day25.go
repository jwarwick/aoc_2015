package main

import (
	"fmt"
)

func main() {
	start := 20151125
	row := 2978
	col := 3083
	part1 := code(start, row, col)
	fmt.Println("Part 1: ", part1)
}

func code(start int, row int, col int) int {
	result := -1
	max_row := 1
	last := start
	for {
		curr_col := 1
		for r := max_row; r >= 1; r-- {
			if r == row && curr_col == col {
				return last
			}
			curr_col++
			v := next(last)
			last = v
		}
		max_row++
	}
	return result
}

func next(value int) int {
	n := value * 252533
	r := n % 33554393
	return int(r)
}
