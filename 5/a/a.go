package main

import (
	five "advent/5"
	"advent/aoc"
	"fmt"
)

func main() {
	input := aoc.DirectedVectors(aoc.Lines("5.input"))

	vectors := make([]five.Vector, len(input))
	for i, raw := range input {
		vectors[i] = five.NewVector(raw)
	}

	plot := make(five.Plot)
	for _, vector := range vectors {
		if !vector.IsStraight() {
			continue
		}

		plot.AddLine(vector.Line())
	}

	fmt.Println(plot.Intersects())
}
