package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rotations []int

	for scanner.Scan() {
		line := scanner.Text()

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

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rotations, nil
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func part1(rotations []int) int {
	cur := 50
	ans := 0

	for _, rot := range rotations {
		cur += rot
		cur = mod(cur, 100)
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
		turns := abs(rot) / 100
		rem := rot % 100
		nxt := cur + rem
		if abs(rem) > 0 && ((nxt <= 0 && cur != 0) || nxt >= 100) {
			ans++
		}
		ans += turns
		cur = mod(nxt, 100)
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
