package main

import (
	"fmt"
	"math"
)

func main() {
	input := 33100000

	part1 := findHouse(input)
	fmt.Println("Part 1: ", part1)
}

func findHouse(target int) int {
	for i := 1; i < 1000000; i++ {
		sum := divisorSum(i)
		sum *= 10
		fmt.Println(i, ":", sum)
		if sum >= target {
			return i
		}
	}
	return -1
}

func divisorSum(value int) int {
	sum := 0
	max := int(math.Sqrt(float64(value)))
	for d := 1; d <= max; d++ {
		if value%d == 0 {
			r := int(value / d)
			sum += d
			if d != r {
				sum += r
			}
		}
	}
	return sum
}
