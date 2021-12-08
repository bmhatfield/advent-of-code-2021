package main

import (
	"advent/aoc"
	"fmt"
)

func main() {
	scrambled := aoc.SignalOutputs(aoc.Lines("input"))

	var uniqs int
	for _, so := range scrambled {
		for _, output := range so[1] {
			switch len(output) {
			case 2, 4, 3, 7:
				uniqs++
			}
		}
	}

	fmt.Println(uniqs)
}
