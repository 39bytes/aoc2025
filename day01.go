package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Rotation struct {
	dir   byte
	count int
}

func getInput(path string) ([]Rotation, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rotations []Rotation

	for scanner.Scan() {
		line := scanner.Text()

		dir := line[0]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}
		rotations = append(rotations, Rotation{
			dir,
			count,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rotations, nil
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func part1(rotations []Rotation) int {
	cur := 50
	ans := 0

	for _, rot := range rotations {
		if rot.dir == 'L' {
			cur -= rot.count
		} else {
			cur += rot.count
		}
		cur = mod(cur, 100)
		if cur == 0 {
			ans++
		}
	}

	return ans
}

func part2(rotations []Rotation) int {
	cur := 50
	ans := 0

	for _, rot := range rotations {
		before := cur
		if rot.dir == 'L' {
			cur -= rot.count

			if before == 0 {
				ans += int(math.Abs(math.Floor(float64(cur)/100.0))) - 1
			} else {
				ans += int(math.Abs(math.Floor(float64(cur) / 100.0)))
			}
			cur = mod(cur, 100)
			if cur == 0 {
				ans++
			}
		} else {
			cur += rot.count
			ans += int(math.Abs(math.Floor(float64(cur) / 100.0)))
			cur = mod(cur, 100)
		}
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
