package main

import (
	"aoc2025/util"
	"fmt"
	"log"
	"strings"
)

func dfs(cur, target string, graph map[string][]string, memo map[string]int) int {
	if cur == target {
		return 1
	}
	key := cur + target
	if x, memoized := memo[key]; memoized {
		return x
	}

	neighbors, _ := graph[cur]
	ans := 0
	for _, nei := range neighbors {
		ans += dfs(nei, target, graph, memo)
	}

	memo[cur+target] = ans

	return ans
}

func part1(graph map[string][]string) int {
	memo := make(map[string]int)
	return dfs("you", "out", graph, memo)
}

func part2(graph map[string][]string) int {
	memo := make(map[string]int)
	return dfs("svr", "fft", graph, memo) * dfs("fft", "dac", graph, memo) * dfs("dac", "out", graph, memo)
}

func main() {
	lines, err := util.Lines("inputs/11/input.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	graph := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		node := parts[0]
		neighbors := strings.Split(parts[1], " ")
		graph[node] = neighbors
	}

	fmt.Printf("Part 1: %d\n", part1(graph))
	fmt.Printf("Part 2: %d", part2(graph))
}
