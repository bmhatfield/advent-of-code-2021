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

func Characters(lines []string) [][]string {
	out := make([][]string, len(lines))

	for i, line := range lines {
		out[i] = strings.Split(line, "")
	}

	return out
}

func Boards(lines []string) ([]string, [][][]string) {
	boards := make([][][]string, 0)
	var drawn []string

	bi := -1
	for i, line := range lines {
		if i == 0 {
			drawn = strings.Split(line, ",")
			continue
		}

		if line == "" {
			bi++
			boards = append(boards, [][]string{})
			continue
		}

		cleaned := strings.ReplaceAll(strings.TrimLeft(line, " "), "  ", " ")
		boardLine := strings.Split(cleaned, " ")

		boards[bi] = append(
			boards[bi],
			boardLine,
		)
	}

	return drawn, boards
}

func DirectedVectors(lines []string) [][2][2]int {
	out := make([][2][2]int, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " -> ")
		start := strings.Split(parts[0], ",")
		end := strings.Split(parts[1], ",")

		out[i] = [2][2]int{
			{
				Int(start[0]),
				Int(start[1]),
			},
			{
				Int(end[0]),
				Int(end[1]),
			},
		}
	}

	return out
}
