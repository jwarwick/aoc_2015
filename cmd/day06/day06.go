package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	part1 := countLights(string(input))
	fmt.Println("Part 1: ", part1)

	part2 := brightLights(string(input))
	fmt.Println("Part 2: ", part2)
}

func brightLights(input string) int {
	lines := splitInput(input)
	areas := make([]area, len(lines))
	for idx, l := range lines {
		areas[idx] = parse_area(l)
	}

	total := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			state := 0
			for _, a := range areas {
				if a.contains(x, y) {
					state = a.bright_fn(state)
				}
			}
			total += state
		}
	}
	return total
}

func countLights(input string) int {
	lines := splitInput(input)
	areas := make([]area, len(lines))
	for idx, l := range lines {
		areas[idx] = parse_area(l)
	}

	total := 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			state := false
			for _, a := range areas {
				if a.contains(x, y) {
					state = a.fn(state)
				}
			}
			if state {
				total++
			}
		}
	}
	return total
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	return lines
}

func parse_area(input string) area {
	var a area
	switch {
	case strings.HasPrefix(input, "turn on"):
		a.fn = turnOn
		a.bright_fn = turnOnBright
	case strings.HasPrefix(input, "turn off"):
		a.fn = turnOff
		a.bright_fn = turnOffBright
	case strings.HasPrefix(input, "toggle"):
		a.fn = toggle
		a.bright_fn = toggleBright
	default:
		panic("No instruction")
	}
	halves := strings.Split(input, " through ")

	min_all := strings.Split(halves[0], " ")
	min_vals := min_all[len(min_all)-1]
	min := strings.Split(min_vals, ",")
	a.min_x = parse_int(min[0])
	a.min_y = parse_int(min[1])

	max := strings.Split(halves[1], ",")
	a.max_x = parse_int(max[0])
	a.max_y = parse_int(max[1])

	return a
}

func parse_int(input string) int {
	i, err := strconv.Atoi(input)
	if nil != err {
		panic("Failed to parse int")
	}
	return i
}

type instruction_func func(bool) bool
type bright_func func(int) int

type area struct {
	min_x, min_y int
	max_x, max_y int
	fn           instruction_func
	bright_fn    bright_func
}

func (a *area) contains(x int, y int) bool {
	return x >= a.min_x && x <= a.max_x &&
		y >= a.min_y && y <= a.max_y
}

func turnOn(val bool) bool {
	return true
}

func turnOff(val bool) bool {
	return false
}

func toggle(val bool) bool {
	return !val
}

func turnOnBright(val int) int {
	return val + 1
}

func turnOffBright(val int) int {
	n := val - 1
	if n < 0 {
		n = 0
	}
	return n
}

func toggleBright(val int) int {
	return val + 2
}
