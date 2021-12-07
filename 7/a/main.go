package main

import (
	"advent/aoc"
	"fmt"
)

func Shift(nums []int, target int) (out int) {
	for _, num := range nums {
		out += aoc.Abs(num - target)
	}

	return out
}

func main() {
	crabs := aoc.SplitInts(aoc.Lines("input"))

	med := aoc.Median(crabs)

	sft := Shift(crabs, med)

	fmt.Println(med, sft)
}
