package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1321131112"
	part1 := lookAndSayMulti(input, 40)
	fmt.Println("Part 1: ", len(part1))

	part2 := lookAndSayMulti(input, 50)
	fmt.Println("Part 2: ", len(part2))
}

func lookAndSayMulti(input string, repeat int) string {
	for i := 0; i < repeat; i++ {
		input = lookAndSay(input)
	}
	return input
}

func lookAndSay(input string) string {
	curr := rune(input[0])
	cnt := 1
	var values []string
	for _, r := range input[1:] {
		if curr == r {
			cnt++
		} else {
			values = append(values, strconv.Itoa(cnt))
			values = append(values, string(curr))
			curr = r
			cnt = 1
		}
	}
	values = append(values, strconv.Itoa(cnt))
	values = append(values, string(curr))
	return strings.Join(values, "")
}
