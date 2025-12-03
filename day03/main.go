package main

import (
	"aoc2025/util"
	"fmt"
	"log"
)

func largest(line []byte, n int) int64 {
	last := -1
	ans := int64(0)

	for i := range n {
		end := len(line) - n + i + 1
		best := last + 1 + util.MaxIndex(line[last+1:end])
		ans = ans*10 + int64(line[best]-'0')
		last = best
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
