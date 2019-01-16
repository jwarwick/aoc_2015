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

	part1 := shortestPath(string(input))
	fmt.Println("Part 1: ", part1)
}

func shortestPath(input string) int {
	lines := splitInput(input)

	distances := make(map[string]map[string]int)
	for _, l := range lines {
		from, to, dist := parseLine(l)
		addDistance(from, to, dist, &distances)
		addDistance(to, from, dist, &distances)
	}

	var cities []string
	for k, _ := range distances {
		cities = append(cities, k)
	}

	p := permutations(cities)
	best := math.MaxInt64
	for _, e := range p {
		d := distance(e, &distances)
		if d < best {
			best = d
		}
	}
	return best
}

func addDistance(a string, b string, d int, m *map[string]map[string]int) {
	if _, prs := (*m)[a]; !prs {
		(*m)[a] = make(map[string]int)
	}
	(*m)[a][b] = d
}

type link struct {
	from, to string
	distance int
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return lines
}

func parseLine(input string) (from string, to string, dist int) {
	halves := strings.Split(input, " = ")
	locs := strings.Split(halves[0], " to ")
	from = locs[0]
	to = locs[1]
	var err error
	dist, err = strconv.Atoi(halves[1])
	if nil != err {
		panic("Not a number when parsing")
	}
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

func distance(path []string, m *map[string]map[string]int) int {
	total := 0
	prev := path[0]
	for _, curr := range path[1:] {
		dists := (*m)[prev]
		total += dists[curr]
		prev = curr
	}
	return total
}
