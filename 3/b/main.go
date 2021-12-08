package main

import (
	"advent/aoc"
	"fmt"
	"strconv"
	"strings"
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
	if b.t < b.f {
		return "1"
	}

	return "0"
}

func report(words [][]string) []bit {
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

	return report
}

type validatorFn func(string, bit) bool

func epsilon(w string, b bit) bool {
	return w == b.epsilon()
}

func gamma(w string, b bit) bool {
	return w == b.gamma()
}

func filter(words [][]string, v validatorFn) []string {
	outset := words

	offset := 0
	for len(outset) > 1 {
		rept := report(outset)
		working := make([][]string, 0, len(outset))

		for _, word := range outset {
			if v(word[offset], rept[offset]) {
				working = append(working, word)
			}
		}

		outset = working
		offset++
	}

	return outset[0]
}

func main() {
	words := aoc.Characters(aoc.Lines("3.input"))

	o2 := filter(words, gamma)
	co2 := filter(words, epsilon)

	g, err := strconv.ParseInt(strings.Join(o2, ""), 2, 64)
	if err != nil {
		panic(err)
	}

	e, err := strconv.ParseInt(strings.Join(co2, ""), 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(g, e, g*e)
}
