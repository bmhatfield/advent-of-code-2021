package main

import (
	"advent/aoc"
	"fmt"
	"strconv"
)

type bit struct {
	t int
	f int
}

func (b bit) gamma() string {
	if b.t >= b.f {
		return "1"
	}

	return "0"
}

func (b bit) epsilon() string {
	if b.t <= b.f {
		return "1"
	}

	return "0"
}

func main() {
	words := aoc.Characters(aoc.Lines("3.input"))

	report := make([]bit, len(words[0]))

	for _, word := range words {
		for j, bit := range word {
			switch bit {
			case "0":
				report[j].f++
			case "1":
				report[j].t++
			default:
				panic("unexpected bit")
			}
		}
	}

	gamma := ""
	epsilon := ""
	for _, bit := range report {
		gamma += bit.gamma()
		epsilon += bit.epsilon()
	}

	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		panic(err)
	}

	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(g * e)
}
