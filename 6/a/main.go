package main

import (
	"advent/aoc"
	"fmt"
)

const days = 80

func main() {
	fish := aoc.SplitInts(aoc.Lines("input"))

	for i := 0; i < days; i++ {
		adds := make([]int, 0)

		for i, f := range fish {
			switch f {
			case 0:
				fish[i] = 6
				adds = append(adds, 8)
			default:
				fish[i] = f - 1
			}
		}

		fish = append(fish, adds...)
	}

	fmt.Println(len(fish))
}
