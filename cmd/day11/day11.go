package main

import (
	"fmt"
)

func main() {
	input := "hepxcrrq"
	part1 := nextPassword(input)
	fmt.Println("Part 1: ", part1)

	part2 := nextPassword(part1)
	fmt.Println("Part 2: ", part2)
}

func nextPassword(input string) string {
	max_attempts := 500000000
	cnt := 0
	p := CreateNewPassword(input)
	p.Increment()
	for !p.Valid() && cnt < max_attempts {
		p.Increment()
		cnt++
	}
	return p.String()
}

type password struct {
	value []int
}

func CreateNewPassword(input string) password {
	p := password{value: make([]int, len(input))}
	for i, r := range input {
		p.value[i] = int(r - 'a')
	}
	return p
}

func (p *password) Increment() {
	l := len((*p).value)
	for i := l - 1; i >= 0; i-- {
		entry := &((*p).value[i])
		*entry = *entry + 1
		if *entry > 25 {
			*entry = 0
		} else {
			return
		}
	}
}

func (p password) String() string {
	runes := make([]rune, len(p.value))
	for i, v := range p.value {
		runes[i] = rune(v + int('a'))
	}
	return string(runes)
}

const i_val int = int('i' - 'a')
const o_val int = int('o' - 'a')
const l_val int = int('l' - 'a')

func (p *password) Valid() bool {
	has_straight := false
	pair_count := 0
	last_pair_idx := -1

	for idx, curr := range (*p).value {
		if idx >= 2 && !has_straight {
			a := (*p).value[idx-2]
			b := (*p).value[idx-1]
			has_straight = (a+1 == b) && (b+1 == curr)
		}

		if curr == i_val || curr == o_val || curr == l_val {
			return false
		}

		if idx >= 1 && idx > last_pair_idx+1 {
			a := (*p).value[idx-1]
			if a == curr {
				pair_count++
				last_pair_idx = idx
			}
		}
	}
	return has_straight && pair_count >= 2
}
