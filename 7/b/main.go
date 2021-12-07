package main

import (
	"advent/aoc"
	"fmt"
	"math"
)

func Shift(nums []int, target int) (out int) {
	for _, num := range nums {
		out += aoc.Summation(
			aoc.Abs(num - target),
		)
	}

	return out
}

func main() {
	crabs := aoc.SplitInts(aoc.Lines("input"))

	avg := aoc.Mean(crabs)

	rvg := int(math.Round(avg))
	min := int(math.Round(avg / 2))
	max := int(math.Round(avg * 2))

	best := Shift(crabs, rvg)
	for i := min; i <= max; i++ {
		sft := Shift(crabs, i)
		if sft < best {
			best = sft
		}
	}

	fmt.Println(best)
}
