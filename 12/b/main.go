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

type history struct {
	r map[string]int

	smallx2 bool
}

func (h *history) visited(node string) bool {
	n := h.r[node]

	switch node {
	case twelve.Start, twelve.End:
		return n > 0
	default:
		if h.smallx2 {
			return n > 0
		}

		return n > 1
	}
}

func (h *history) visit(node string) *history {
	out := make(map[string]int)
	for k, v := range h.r {
		out[k] = v
	}

	out[node]++

	return &history{
		r:       out,
		smallx2: h.smallx2 || out[node] > 1,
	}
}

func descend(idx twelve.Index, hist *history, left string) [][]string {
	if small(left) {
		hist = hist.visit(left)
	}

	paths := make([][]string, 0)

	if len(idx[left]) == 0 {
		return [][]string{{left}}
	}

	for right := range idx[left] {
		if small(right) && hist.visited(right) {
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
	edges := aoc.Edges(aoc.Lines("input"))

	idx := twelve.NewIndex()

	for _, edge := range edges {
		idx.Insert(edge[0], edge[1])
	}

	paths := descend(idx, &history{r: map[string]int{}}, twelve.Start)

	fmt.Println(len(paths))
}
