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

	part1 := minMD5(Input)
	fmt.Println("Part 1: ", part1)

	// part2 := housesDeliveredTo(string(input), 2)
	// fmt.Println("Part 2: ", part2)
}

func minMD5(input string) int {
	for i := 0; i < 2500000; i++ {
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

func check(input []byte) bool {
	return input[0] == 0 &&
		input[1] == 0 &&
		input[2] <= 0x0f
}
