package main

import (
	"aoc2025/util"
	"fmt"
	"log"
	"strconv"
)

func getInput(path string) ([]int, error) {
	var rotations []int
	lines, err := util.Lines(path)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		dir := line[0]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}
		if dir == 'L' {
			count = -count
		}
		rotations = append(rotations, count)
	}

	return rotations, nil
}

func part1(rotations []int) int {
	cur := 50
	ans := 0

	for _, rot := range rotations {
		cur += rot
		cur = util.Mod(cur, 100)
		if cur == 0 {
			ans++
		}
	}

	return ans
}

func part2(rotations []int) int {
	cur := 50
	ans := 0

	for _, rot := range rotations {
		turns := util.Abs(rot) / 100
		rem := rot % 100
		nxt := cur + rem
		if util.Abs(rem) > 0 && ((nxt <= 0 && cur != 0) || nxt >= 100) {
			ans++
		}
		ans += turns
		cur = util.Mod(nxt, 100)
	}

	return ans
}

func main() {
	rotations, err := getInput("inputs/01/input.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Printf("Part 1: %d\n", part1(rotations))
	fmt.Printf("Part 2: %d\n", part2(rotations))
}
