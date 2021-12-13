package main

import (
	thirteen "advent/13"
	"advent/aoc"
	"fmt"
)

func visible(grid [][]bool) int {
	var count int
	for _, row := range grid {
		for _, col := range row {
			if col {
				count++
			}
		}
	}

	return count
}

func main() {
	coords, instrs := aoc.ActivationCode(aoc.Lines("input"))

	height, width := thirteen.Dimensions(coords)

	grid := thirteen.MkGrid(height, width)
	grid = thirteen.Plot(grid, coords)

	folded := thirteen.Fold(grid, instrs[0])
	fmt.Println(visible(folded))
}
