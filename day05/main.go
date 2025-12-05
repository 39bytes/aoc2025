package main

import (
	"aoc2025/util"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Interval struct {
	start int
	end   int
}

func part1(intervals []Interval, queries []int) int {
	ans := int(0)

	for _, x := range queries {
		fresh := false
		for _, r := range intervals {
			if r.start <= x && x <= r.end {
				fresh = true
				break
			}
		}

		if fresh {
			ans++
		}
	}

	return ans
}

func part2(intervals []Interval) int {
	slices.SortFunc(intervals, func(a, b Interval) int {
		if a.start == b.start {
			return a.end - b.end
		} else {
			return a.start - b.start
		}
	})

	ans := 0
	start := intervals[0].start
	end := intervals[0].end

	for i := 1; i < len(intervals); i++ {
		if intervals[i].start <= end {
			end = max(intervals[i].end, end)
		} else {
			ans += end - start + 1
			start = intervals[i].start
			end = intervals[i].end
		}
	}

	ans += end - start + 1
	return ans
}

func main() {
	lines, err := util.Lines("inputs/05/input.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	var intervals []Interval
	i := 0
	for lines[i] != "" {
		parts := strings.Split(lines[i], "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		intervals = append(intervals, Interval{start, end})
		i++
	}

	var queries []int
	i++
	for i < len(lines) {
		num, _ := strconv.Atoi(lines[i])
		queries = append(queries, num)
		i++
	}

	fmt.Printf("Part 1: %d\n", part1(intervals, queries))
	fmt.Printf("Part 2: %d\n", part2(intervals))
}
