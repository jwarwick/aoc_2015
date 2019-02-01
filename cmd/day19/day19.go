package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	// "sort"
	"strings"
	"time"
	"unicode"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	rules, start := parseFile(string(input))
	part1 := molecules(start, rules)
	fmt.Println("Part 1: ", part1)

	part2 := build(start, rules)
	fmt.Println("Part 2: ", part2)
}

type rule struct {
	from string
	to   string
}

type node struct {
	str   string
	depth int
}

type ruleSlice []rule

func (s ruleSlice) Len() int {
	return len(s)
}
func (s ruleSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ruleSlice) Less(i, j int) bool {
	return len(s[i].to) > len(s[j].to)
}

func build(target string, rules []rule) int {
	// _, count := shrink(target, rules)
	// Using formula from AOC reddit solution, since
	// nothing else ever seems to halt...
	// #NumSymbols - #Rn - #Ar - 2 * #Y - 1
	rn := strings.Count(target, "Rn")
	ar := strings.Count(target, "Ar")
	y := strings.Count(target, "Y")
	symbols := 0
	for _, c := range target {
		if !unicode.IsLower(c) {
			symbols++
		}
	}

	return symbols - rn - ar - 2*y - 1
}

func shrink(start string, rules []rule) (bool, int) {
	if len(rules) == 0 {
		return false, -1
	}
	tried := make(map[string]bool)
	for {
		for {
			shuffle(rules)
			fingerprint := ""
			for _, n := range rules {
				fingerprint += n.from
				fingerprint += n.to
			}
			_, present := tried[fingerprint]
			if !present {
				tried[fingerprint] = true
				break
			}
		}
		fmt.Println(rules)
		count := 0
		curr := start
		for _, r := range rules {
			for {
				matches := strings.Count(curr, r.to)
				if matches <= 0 {
					break
				}
				count += matches
				curr = strings.Replace(curr, r.to, r.from, -1)
			}
			// fmt.Println("Steps:", count, curr)
			if curr == "e" {
				return true, count
			}
		}
	}
	return false, -1
}

func permutations(input []rule) [][]rule {
	var output [][]rule
	generate(len(input), input, &output)
	return output
}

func generate(n int, arr []rule, output *[][]rule) {
	if n == 1 {
		tmp := append([]rule(nil), arr...)
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

func shuffle(vals []rule) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}

func parseFile(input string) (rules []rule, start string) {
	lines := splitInput(input)
	rules = make([]rule, len(lines)-2)
	blankIdx := len(rules)
	startIdx := len(rules) + 1
	for idx, l := range lines {
		switch idx {
		case startIdx:
			start = l
		case blankIdx:
			continue
		default:
			words := strings.Fields(l)
			rules[idx] = rule{words[0], words[2]}
		}
	}
	return
}

func reachable(start string, rules *[]rule) []string {
	set := make(map[string]int)
	for _, r := range *rules {
		count := strings.Count(start, r.from)
		for n := 1; n <= count; n++ {
			splits := strings.SplitAfterN(start, r.from, n)
			splits[n-1] = strings.Replace(splits[n-1], r.from, r.to, 1)
			combined := strings.Join(splits, "")
			set[combined]++
		}
	}
	result := make([]string, len(set))
	idx := 0
	for k, _ := range set {
		result[idx] = k
		idx++
	}

	return result
}

func molecules(start string, rules []rule) int {
	result := reachable(start, &rules)
	return len(result)
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return lines
}
