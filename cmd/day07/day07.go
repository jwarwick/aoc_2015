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

	part1 := simulate(string(input), "a", make(map[string]uint16))
	fmt.Println("Part 1: ", part1["a"])

	a := part1["a"]
	part2 := simulate(string(input), "a", map[string]uint16{"b": a})
	fmt.Println("Part 2: ", part2["a"])
}

var gates map[string]gate
var wires map[string]uint16

func simulate(input string, target string, init_wires map[string]uint16) map[string]uint16 {
	gates = make(map[string]gate)
	wires = init_wires

	lines := splitInput(input)

	for _, l := range lines {
		g := parseGate(l)
		gates[g.output] = g
	}

	g := gates[target]
	g.evaluate()

	return wires
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	return lines
}

type gate struct {
	input1 string
	input2 string
	output string
	fn     func(string, string) uint16
}

func parseGate(input string) gate {
	var g gate

	halves := strings.Split(input, " -> ")
	g.output = strings.TrimSpace(halves[1])
	cmd := strings.TrimSpace(halves[0])
	switch {
	case strings.Contains(cmd, "AND"):
		args := strings.Split(cmd, " AND ")
		g.input1 = args[0]
		g.input2 = args[1]
		g.fn = and
	case strings.Contains(cmd, "OR"):
		args := strings.Split(cmd, " OR ")
		g.input1 = args[0]
		g.input2 = args[1]
		g.fn = or
	case strings.Contains(cmd, "NOT"):
		args := strings.Split(cmd, "NOT ")
		g.input1 = args[1]
		g.fn = not
	case strings.Contains(cmd, "LSHIFT"):
		args := strings.Split(cmd, " LSHIFT ")
		g.input1 = args[0]
		g.input2 = args[1]
		g.fn = lshift
	case strings.Contains(cmd, "RSHIFT"):
		args := strings.Split(cmd, " RSHIFT ")
		g.input1 = args[0]
		g.input2 = args[1]
		g.fn = rshift
	default:
		g.input1 = cmd
		g.fn = constant
	}

	return g
}

func (g *gate) evaluate() uint16 {
	_, present := wires[g.output]
	if !present {
		result := g.fn(g.input1, g.input2)
		wires[g.output] = result
	}
	return wires[g.output]
}

func getValue(input string) uint16 {
	i, err := strconv.Atoi(input)
	if nil == err {
		return uint16(i)
	}

	_, present := wires[input]
	if !present {
		g := gates[input]
		g.evaluate()
	}
	return wires[input]
}

func parseInt(input string) int {
	i, err := strconv.Atoi(input)
	if nil != err {
		fmt.Println("Not an int: ", input)
		panic("Failed to parse int")
	}
	return i
}

func constant(input string, _ string) uint16 {
	val := getValue(input)
	return val
}

func not(input string, _ string) uint16 {
	val := getValue(input)
	return ^val
}

func and(input1 string, input2 string) uint16 {
	v1 := getValue(input1)
	v2 := getValue(input2)
	return v1 & v2
}

func or(input1 string, input2 string) uint16 {
	v1 := getValue(input1)
	v2 := getValue(input2)
	return v1 | v2
}

func lshift(input1 string, input2 string) uint16 {
	v1 := getValue(input1)
	v2 := parseInt(input2)
	return v1 << uint(v2)
}

func rshift(input1 string, input2 string) uint16 {
	v1 := getValue(input1)
	v2 := parseInt(input2)
	return v1 >> uint(v2)
}
