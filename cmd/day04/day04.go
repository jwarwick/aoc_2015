package main

import (
	"crypto/md5"
	// "encoding/hex"
	"fmt"
	"io"
	"strconv"
)

func main() {
	const Input = "yzbqklnj"

	part1 := minMD5(Input, part1Check)
	fmt.Println("Part 1: ", part1)

	part2 := minMD5(Input, part2Check)
	fmt.Println("Part 2: ", part2)
}

func minMD5(input string, check check_fun) int {
	for i := 0; i < 25000000; i++ {
		result := hash(input, i)
		if check(result) {
			return i
		}
	}
	return -1
}

func hash(input string, cnt int) []byte {
	h := md5.New()
	io.WriteString(h, input)
	io.WriteString(h, strconv.Itoa(cnt))
	return h.Sum(nil)
}

type check_fun func([]byte) bool

func part1Check(input []byte) bool {
	return input[0] == 0 &&
		input[1] == 0 &&
		input[2] <= 0x0f
}

func part2Check(input []byte) bool {
	return input[0] == 0 &&
		input[1] == 0 &&
		input[2] == 0
}
