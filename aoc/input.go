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

type Pair struct {
	Key   string
	Value int
}

func Pairs(lines []string) []Pair {
	out := make([]Pair, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")

		key := parts[0]
		valStr := parts[1]

		n, err := strconv.Atoi(valStr)
		if err != nil {
			panic(err)
		}
		out[i] = Pair{
			Key:   key,
			Value: n,
		}
	}

	return out
}
