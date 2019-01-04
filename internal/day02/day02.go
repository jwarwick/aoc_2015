package day02

import (
	// "fmt"
	"strconv"
	"strings"
)

type computation func(string) int

func WrappingPaperSum(input string) int {
	return sum(input, PaperNeeded)
}

func RibbonSum(input string) int {
	return sum(input, RibbonNeeded)
}

func sum(input string, fn computation) int {
	total := 0
	lines := splitInput(input)
	for _, l := range lines {
		total += fn(l)
	}
	return total
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	return lines
}

func getSides(input string) (int, int, int) {
	s := strings.Split(input, "x")
	l := convertString(s[0])
	w := convertString(s[1])
	h := convertString(s[2])
	return l, w, h
}

func RibbonNeeded(input string) int {
	l, w, h := getSides(input)
	sides := make([]int, 3)
	sides[0] = 2*l + 2*w
	sides[1] = 2*w + 2*h
	sides[2] = 2*h + 2*l
	min := minSide(sides)
	volume := l * w * h
	return min + volume
}

func PaperNeeded(input string) int {
	l, w, h := getSides(input)
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
