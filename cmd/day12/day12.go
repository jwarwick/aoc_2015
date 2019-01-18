package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	part1 := stringSum(string(input))
	fmt.Println("Part 1: ", part1)

	part2 := notRedStringSum(string(input))
	fmt.Println("Part 2: ", part2)
}

func stringSum(input string) int {
	var numbers = regexp.MustCompile(`[-]?\d{1,}`)
	matches := numbers.FindAllString(input, -1)
	total := 0
	for _, s := range matches {
		val, err := strconv.Atoi(s)
		if nil != err {
			fmt.Println("Invalid: ", val)
			panic("Not a number")
		}
		total += val
	}
	return total
}

func notRedStringSum(input string) int {
	var i interface{}
	err := json.Unmarshal([]byte(input), &i)
	if err != nil {
		panic("Failed to parse json")
	}
	return parseEntry(i)
}

func parseMap(m map[string]interface{}) int {
	for _, v := range m {
		if v == "red" {
			return 0
		}
	}

	total := 0
	for _, v := range m {
		total += parseEntry(v)
	}
	return total
}

func parseArray(a []interface{}) int {
	total := 0
	for _, v := range a {
		total += parseEntry(v)
	}
	return total
}

func parseEntry(v interface{}) int {
	total := 0
	switch v.(type) {
	case string:
	case float64:
		total += int(v.(float64))
	case map[string]interface{}:
		total += parseMap(v.(map[string]interface{}))
	case []interface{}:
		total += parseArray(v.([]interface{}))
	default:
		fmt.Println("Entry: ", v)
		panic("Unknown type")
	}
	return total
}
