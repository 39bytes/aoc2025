package main

import (
	"aoc2025/util"
	"fmt"
	"log"
	"strings"
)

type Coord struct {
	i int
	j int
}

func dfs(i, j int, grid []string, splits map[Coord]struct{}) int {
	if i >= len(grid) {
		return 0
	}

	if grid[i][j] == '.' || grid[i][j] == 'S' {
		return dfs(i+1, j, grid, splits)
	}

	_, left_exists := splits[Coord{i, j - 1}]
	_, right_exists := splits[Coord{i, j + 1}]
	if left_exists && right_exists {
		return 0
	}
	ans := 1
	if !left_exists {
		splits[Coord{i, j - 1}] = struct{}{}
		ans += dfs(i, j-1, grid, splits)
	}
	if !right_exists {
		splits[Coord{i, j + 1}] = struct{}{}
		ans += dfs(i, j+1, grid, splits)
	}
	return ans
}

func part1(start_col int, grid []string) int {
	splits := make(map[Coord]struct{})
	return dfs(0, start_col, grid, splits)
}

func dfs2(i, j int, grid []string, memo map[Coord]int) int {
	if i >= len(grid) {
		return 1
	}
	x, ok := memo[Coord{i, j}]
	if ok {
		return x
	}

	var ans int
	if grid[i][j] == '.' || grid[i][j] == 'S' {
		ans = dfs2(i+1, j, grid, memo)
	} else {
		ans = dfs2(i, j-1, grid, memo) + dfs2(i, j+1, grid, memo)
	}
	memo[Coord{i, j}] = ans
	return ans
}

func part2(start_col int, grid []string) int {
	memo := make(map[Coord]int)
	return dfs2(0, start_col, grid, memo)
}

func main() {
	lines, err := util.Lines("inputs/07/input.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	start_col := strings.Index(lines[0], "S")
	fmt.Printf("Part 1: %d\n", part1(start_col, lines))
	fmt.Printf("Part 2: %d\n", part2(start_col, lines))
}
