package main

import (
	thirteen "advent/13"
	"advent/aoc"
	"fmt"
)

func main() {
	coords, instrs := aoc.ActivationCode(aoc.Lines("input"))

	height, width := thirteen.Dimensions(coords)

	grid := thirteen.MkGrid(height, width)
	grid = thirteen.Plot(grid, coords)

	folded := grid
	for _, instr := range instrs {
		folded = thirteen.Fold(folded, instr)
	}

	fmt.Println(thirteen.SGrid(folded))
}
