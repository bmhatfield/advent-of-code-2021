package main

import (
	"advent/aoc"
	"fmt"
)

type window [3]int

func (w window) sum() int {
	return w[0] + w[1] + w[2]
}

func main() {
	soundings := aoc.Ints(aoc.Lines("1.input"))

	increases := 0
	leading := window{}
	following := window{}

	for i, n := range soundings {
		if i+1 > len(soundings)-1 {
			break
		}

		li := i % 3
		fi := (i + 1) % 3

		leading[li] = n
		following[fi] = soundings[i+1]

		if i > 1 {
			a, b := leading.sum(), following.sum()

			if b > a {
				increases++
			}
		}
	}

	fmt.Println(increases)
}
