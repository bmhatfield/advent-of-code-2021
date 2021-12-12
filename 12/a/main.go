package main

import (
	twelve "advent/12"
	"advent/aoc"
	"fmt"
	"strings"
)

func small(node string) bool {
	return strings.ToLower(node) == node
}

func visited(hist map[string]int, node string) bool {
	n := hist[node]

	switch node {
	case twelve.Start, twelve.End:
		return n > 0
	default:
		return n > 0
	}
}

func visit(hist map[string]int, node string) map[string]int {
	out := make(map[string]int)
	for k, v := range hist {
		out[k] = v
	}

	out[node]++

	return out
}

func descend(idx twelve.Index, hist map[string]int, left string) [][]string {
	if small(left) {
		hist = visit(hist, left)
	}

	paths := make([][]string, 0)

	if len(idx[left]) == 0 {
		return [][]string{{left}}
	}

	for right := range idx[left] {
		if small(right) && visited(hist, right) {
			continue
		}

		for _, step := range descend(idx, hist, right) {
			paths = append(paths, append([]string{left}, step...))
		}
	}

	return paths
}

func pathjoin(ps [][]string) string {
	var out string
	for _, paths := range ps {
		out += fmt.Sprintln(strings.Join(paths, ","))
	}
	return out
}

func main() {
	edges := aoc.Edges(aoc.Lines("test"))

	idx := twelve.NewIndex()

	for _, edge := range edges {
		idx.Insert(edge[0], edge[1])
	}

	paths := descend(idx, make(map[string]int), twelve.Start)

	fmt.Println(pathjoin(paths), len(paths))
}
