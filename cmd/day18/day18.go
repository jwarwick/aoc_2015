package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	part1 := simulate(string(input), 100, false)
	fmt.Println("Part 1: ", part1)

	part2 := simulate(string(input), 100, true)
	fmt.Println("Part 2: ", part2)
}

func simulate(input string, steps int, cornersOn bool) int {
	b := createBoard(input)
	if cornersOn {
		b.cornersOn()
	}
	// fmt.Println("Initial State:")
	// b.display()
	for i := 0; i < steps; i++ {
		b.step()
		if cornersOn {
			b.cornersOn()
		}
		// fmt.Println("After Step ", i+1)
		// b.display()
	}
	return b.count()
}

type board struct {
	size int
	grid [][]bool
}

func (b *board) cornersOn() {
	b.grid[0][0] = true
	b.grid[0][b.size-1] = true
	b.grid[b.size-1][0] = true
	b.grid[b.size-1][b.size-1] = true
}

func (b *board) count() int {
	count := 0
	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			if b.grid[r][c] {
				count++
			}
		}
	}
	return count
}

func (b *board) step() {
	new_grid := make([][]bool, b.size)
	for i := range new_grid {
		new_grid[i] = make([]bool, b.size)
	}

	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			state := b.grid[r][c]
			cnt := b.neighbors(r, c)
			if state {
				switch cnt {
				case 2, 3:
					new_grid[r][c] = true
				default:
					new_grid[r][c] = false
				}
			} else {
				switch cnt {
				case 3:
					new_grid[r][c] = true
				default:
					new_grid[r][c] = false
				}
			}
		}
	}

	b.grid = new_grid
}

func (b *board) neighbors(baseRow int, baseCol int) int {
	cnt := 0
	for r := baseRow - 1; r <= baseRow+1; r++ {
		for c := baseCol - 1; c <= baseCol+1; c++ {
			if r < 0 || c < 0 || r >= b.size || c >= b.size || (r == baseRow && c == baseCol) {
				continue
			}
			if b.grid[r][c] {

				cnt++
			}
		}
	}

	return cnt
}

func (b *board) display() {
	for r := 0; r < b.size; r++ {
		for c := 0; c < b.size; c++ {
			if b.grid[r][c] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func createBoard(input string) board {
	lines := splitInput(input)
	size := len(lines)
	grid := make([][]bool, size)
	for i := range grid {
		grid[i] = make([]bool, size)
	}
	b := board{size, grid}
	for row, l := range lines {
		for col, c := range l {
			grid[row][col] = c == '#'
		}
	}
	return b
}

func splitInput(input string) []string {
	trimmed := strings.TrimSpace(input)
	lines := strings.Split(trimmed, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return lines
}
