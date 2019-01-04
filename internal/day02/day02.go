package day02

import (
	// "fmt"
	"strconv"
	"strings"
)

func WrappingPaperSum(input string) int {
	total := 0
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	for _, l := range lines {
		total += PaperNeeded(l)
	}
	return total
}

func PaperNeeded(input string) int {
	s := strings.Split(input, "x")
	l := convertString(s[0])
	w := convertString(s[1])
	h := convertString(s[2])
	sides := make([]int, 3)
	sides[0] = l * w
	sides[1] = w * h
	sides[2] = h * l
	min := minSide(sides)
	total := (2 * sides[0]) + (2 * sides[1]) + (2 * sides[2]) + min
	return total
}

func minSide(sides []int) int {
	min := sides[0]
	for _, s := range sides {
		if s < min {
			min = s
		}
	}
	return min
}

func convertString(input string) int {
	v, err := strconv.Atoi(input)
	if nil != err {
		panic(err)
	}
	return v
}
