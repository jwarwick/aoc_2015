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

	fmt.Println(cities)
	fmt.Println(distances)
	return 0
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

// procedure generate(n : integer, A : array of any):
//     if n = 1 then
//           output(A)
//     else
//         for i := 0; i < n - 1; i += 1 do
//             generate(n - 1, A)
//             if n is even then
//                 swap(A[i], A[n-1])
//             else
//                 swap(A[0], A[n-1])
//             end if
//         end for
//         generate(n - 1, A)
//     end if
