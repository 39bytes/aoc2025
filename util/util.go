package util

import (
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
