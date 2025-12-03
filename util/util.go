package util

import (
	"cmp"
	"os"
	"strings"
)

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func IntPow(x int64, p int64) int64 {
	res := int64(1)
	for range p {
		res *= x
	}
	return res
}

func Lines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}

func Map[S ~[]E, E any, T any](s S, f func(x E) T) []T {
	mapped := make([]T, len(s))
	for i, x := range s {
		mapped[i] = f(x)
	}
	return mapped
}

func Filter[S ~[]E, E any](s S, f func(x E) bool) []E {
	var filtered []E
	for _, x := range s {
		if f(x) {
			filtered = append(filtered, x)
		}
	}
	return filtered
}

func Sum[T ~int | ~int64](xs []T) int64 {
	sum := int64(0)
	for _, x := range xs {
		sum += int64(x)
	}
	return sum
}

func MaxIndex[S ~[]E, E cmp.Ordered](s S) int {
	max := 0
	for i := range s {
		if s[i] > s[max] {
			max = i
		}
	}
	return max
}
