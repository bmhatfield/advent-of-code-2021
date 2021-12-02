package aoc

import (
	"os"
	"strconv"
	"strings"
)

func Lines(path string) []string {
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	file := strings.TrimSuffix(string(f), "\n")

	return strings.Split(file, "\n")
}

func Ints(lines []string) []int {
	out := make([]int, len(lines))
	for i, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		out[i] = n
	}

	return out
}
