package main

import (
	"advent/aoc"
	"fmt"
	"time"
)

func bounds(counts map[string]int) (int, int) {
	var min, max int
	for _, v := range counts {
		if v < min || min == 0 {
			min = v
		}

		if v > max {
			max = v
		}
	}
	return min, max
}

func main() {
	template, rules := aoc.PolymerRules(aoc.Lines("input"))

	ts := time.Now()
	input := template
	for step := 0; step < 10; step++ {
		fmt.Println(step, time.Since(ts))
		ts = time.Now()

		until := len(input) - 1
		final := until - 1

		var work string
		for i := 0; i < until; i++ {
			p := input[i : i+2]

			ins := rules[p]
			work += string(p[0]) + ins

			if i == final {
				work += string(p[1])
			}
		}

		input = work
	}

	counts := make(map[string]int)
	for _, v := range input {
		counts[string(v)]++
	}

	fmt.Println(bounds(counts))
}
