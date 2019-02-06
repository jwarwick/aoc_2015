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

	part1 := simulate(string(input), 0)
	fmt.Println("Part 1: ", part1)

	part2 := simulate(string(input), 1)
	fmt.Println("Part 2: ", part2)
}

func simulate(input string, regA int) int {
	p := createProgram(input)
	p.registers[0] = regA
	p.run()
	return p.registers[1]
}

type instructionType int

const (
	hlf instructionType = iota
	tpl
	inc
	jmp
	jie
	jio
)

type instruction struct {
	inst instructionType
	arg1 int
	arg2 int
}

func (i instruction) toString() string {
	op := lookupOpCode(i.inst)
	return fmt.Sprintf("%s %d %d", op, i.arg1, i.arg2)
}

type program struct {
	pc           int
	instructions []instruction
	registers    []int
}

func createProgram(input string) program {
	lines := splitInput(input)
	size := len(lines)
	instructions := make([]instruction, size)
	for i, l := range lines {
		instructions[i] = parseLine(l)
	}
	regs := make([]int, 2)
	p := program{pc: 0, instructions: instructions, registers: regs}
	return p
}

func (p *program) run() {
	for p.pc >= 0 && p.pc < len(p.instructions) {
		i := p.instructions[p.pc]
		switch i.inst {
		case hlf:
			p.registers[i.arg1] = p.registers[i.arg1] / 2
			p.pc++
		case tpl:
			p.registers[i.arg1] *= 3
			p.pc++
		case inc:
			p.registers[i.arg1] += 1
			p.pc++
		case jmp:
			p.pc += i.arg1
		case jie:
			if p.registers[i.arg1]%2 == 0 {
				p.pc += i.arg2
			} else {
				p.pc++
			}
		case jio:
			if p.registers[i.arg1] == 1 {
				p.pc += i.arg2
			} else {
				p.pc++
			}
		default:
			panic("Unknown instruction")
		}
	}
}

func parseLine(input string) instruction {
	fields := strings.Fields(input)
	it := lookupType(fields[0])
	a1, a2 := parseArgs(it, fields)
	return instruction{inst: it, arg1: a1, arg2: a2}
}

func parseArgs(iType instructionType, fields []string) (a1 int, a2 int) {
	switch iType {
	case hlf, tpl, inc:
		a1 = parseRegister(fields[1])
	case jmp:
		a1 = parseInt(fields[1])
	case jie, jio:
		a1 = parseRegister(fields[1])
		a2 = parseInt(fields[2])
	default:
		panic("Unknown instruction")
	}
	return
}

func lookupType(input string) instructionType {
	m := map[string]instructionType{"hlf": hlf, "tpl": tpl, "inc": inc, "jmp": jmp, "jie": jie, "jio": jio}
	t, present := m[input]
	if !present {
		panic("Unknown instruction")
	}
	return t
}

func lookupOpCode(input instructionType) string {
	m := map[instructionType]string{hlf: "hlf", tpl: "tpl", inc: "inc", jmp: "jmp", jie: "jie", jio: "jio"}
	t, present := m[input]
	if !present {
		panic("Unknown instruction")
	}
	return t
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return lines
}

func parseRegister(input string) int {
	trimmed := strings.Trim(input, ",")
	switch trimmed {
	case "a":
		return 0
	case "b":
		return 1
	default:
		panic("Unknown register")
	}
}

func parseInt(input string) int {
	trimmed := strings.Trim(input, ",")
	i, err := strconv.Atoi(trimmed)
	if nil != err {
		panic("Not a number when parsing")
	}
	return i
}
