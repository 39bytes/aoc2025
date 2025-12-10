package main

import (
	"aoc2025/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func part1(pts []Point) int {
	ans := 0
	for i := range pts {
		for j := i + 1; j < len(pts); j++ {
			p1 := pts[i]
			p2 := pts[j]
			w := util.Abs(p2.x-p1.x) + 1
			h := util.Abs(p2.y-p1.y) + 1
			ans = max(ans, w*h)
		}
	}

	return ans
}

func intersect(r1, r2, a1, a2 Point) bool {
	tl_x, br_x := min(r1.x, r2.x), max(r1.x, r2.x)
	tl_y, br_y := min(r1.y, r2.y), max(r1.y, r2.y)
	x1, x2 := min(a1.x, a2.x), max(a1.x, a2.x)
	y1, y2 := min(a1.y, a2.y), max(a1.y, a2.y)

	return tl_x < x2 && br_x > x1 && tl_y < y2 && br_y > y1
}

func valid(r1, r2 Point, vertices []Point) bool {
	for i, a := range vertices {
		b := vertices[(i+1)%len(vertices)]

		if intersect(r1, r2, a, b) {
			return false
		}
	}
	return true
}

func part2(pts []Point) int {
	ans := 0
	for i := range pts {
		for j := i + 1; j < len(pts); j++ {
			p1 := pts[i]
			p2 := pts[j]

			if valid(p1, p2, pts) {
				w := util.Abs(p2.x-p1.x) + 1
				h := util.Abs(p2.y-p1.y) + 1
				ans = max(ans, w*h)
			}
		}
	}

	return ans
}

func main() {
	lines, err := util.Lines("inputs/09/input.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	var pts []Point
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		pts = append(pts, Point{x, y})
	}

	fmt.Printf("Part 1: %d\n", part1(pts))
	fmt.Printf("Part 2: %d\n", part2(pts))
}
