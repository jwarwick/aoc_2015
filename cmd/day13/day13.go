package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	part1 := optimalTotal(string(input))
	fmt.Println("Part 1: ", part1)
}

func optimalTotal(input string) int {
	costs := buildCosts(input)
	diners := make([]string, 0)
	for k, _ := range costs {
		diners = append(diners, k)
	}
	perms := permutations(diners)
	best := math.MinInt64
	for _, p := range perms {
		cost := computeCost(p, &costs)
		if cost > best {
			best = cost
		}
	}
	return best
}

func computeCost(diners []string, costs *costTable) int {
	total := 0
	for i, d := range diners {
		neighbor_idx := (i + 1) % len(diners)
		neighbor := diners[neighbor_idx]
		total += (*costs)[d][neighbor]
		total += (*costs)[neighbor][d]
	}
	return total
}

type costTable map[string]map[string]int

func buildCosts(input string) costTable {
	costs := make(costTable)
	lines := splitInput(input)
	for _, l := range lines {
		from, to, val := parseLine(l)
		_, present := costs[from]
		if !present {
			costs[from] = make(map[string]int)
		}
		costs[from][to] = val
	}
	return costs
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return lines
}

func parseLine(input string) (from string, to string, cost int) {
	words := strings.Split(input, " ")
	from = words[0]
	to = strings.Trim(words[10], ".")
	var err error
	cost, err = strconv.Atoi(words[3])
	if nil != err {
		panic("Not a number when parsing")
	}
	mult := 1
	if words[2] == "lose" {
		mult = -1
	}
	cost *= mult
	return
}

func permutations(input []string) [][]string {
	var output [][]string
	generate(len(input), input, &output)
	return output
}

func generate(n int, arr []string, output *[][]string) {
	if n == 1 {
		tmp := append([]string(nil), arr...)
		*output = append(*output, tmp)
	} else {
		for i := 0; i < n-1; i++ {
			generate(n-1, arr, output)
			if n%2 == 0 {
				tmp := arr[i]
				arr[i] = arr[n-1]
				arr[n-1] = tmp
			} else {
				tmp := arr[0]
				arr[0] = arr[n-1]
				arr[n-1] = tmp
			}
		}
		generate(n-1, arr, output)
	}
}
