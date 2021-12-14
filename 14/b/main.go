package main

import (
	"advent/aoc"
	"fmt"
)

func pairs(r map[string]string, a string) (string, string) {
	v := r[a]

	return string(a[0]) + v, v + string(a[1])
}

func main() {
	template, rules := aoc.PolymerRules(aoc.Lines("input"))

	seed := make(map[string]uint)
	for i := range template {
		if i == len(template)-1 {
			break
		}

		p := template[i : i+2]
		seed[p]++
	}

	input := seed
	for step := 0; step < 40; step++ {
		work := make(map[string]uint)

		for k, v := range input {
			l, r := pairs(rules, k)
			work[l] += v
			work[r] += v
		}

		input = work
	}

	counts := make(map[string]uint)
	for k, v := range input {
		counts[string(k[0])] += v
		counts[string(k[1])] += v
	}

	var min, max uint
	for k, v := range counts {
		count := v/2 + v%2
		fmt.Println(k, count)

		if count > max {
			max = count
		}

		if count < min || min == 0 {
			min = count
		}
	}

	fmt.Println(max - min)
}
