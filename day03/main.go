package main

import (
	"aoc2025/util"
	"fmt"
	"log"
)

func largest(line []byte, n int) int64 {
	var stack []byte

	for i := range len(line) {
		for len(stack) > 0 && n-len(stack) < len(line)-i && stack[len(stack)-1] < line[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < n {
			stack = append(stack, line[i])
		}
	}

	ans := int64(0)
	for _, x := range stack {
		ans = ans*10 + int64(x-'0')
	}
	return ans
}

func main() {
	lines, err := util.Lines("inputs/03/input.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	part1 := util.Sum(util.Map(lines, func(s string) int64 { return largest([]byte(s), 2) }))
	part2 := util.Sum(util.Map(lines, func(s string) int64 { return largest([]byte(s), 12) }))

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
