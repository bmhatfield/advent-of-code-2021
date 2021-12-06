package main

import (
	"advent/aoc"
	"fmt"
)

const days = 256

func main() {
	init := aoc.SplitInts(aoc.Lines("input"))

	fish := make([]int, 9)
	for _, f := range init {
		fish[f]++
	}

	for i := 0; i < days; i++ {
		z := fish[0]
		fish = append(fish[1:], z)
		fish[6] += z
	}

	var sum int
	for _, num := range fish {
		sum += num
	}

	fmt.Println(sum)
}
