package main

import (
	ten "advent/10"
	"advent/aoc"
	"fmt"
	"sort"
	"strings"
)

func main() {
	lines := aoc.Lines("input")

	scores := make([]int, 0, len(lines))

	for _, line := range lines {
		s := ten.Stack{}

		var corrupt bool
		for _, ch := range strings.Split(line, "") {
			switch ch {
			case "(":
				s = s.Push(")")
			case "{":
				s = s.Push("}")
			case "[":
				s = s.Push("]")
			case "<":
				s = s.Push(">")

			case ")", "}", "]", ">":
				var expected string
				expected, s = s.Pop()

				if ch != expected {
					corrupt = true
				}
			}

			if corrupt {
				break
			}
		}

		if !corrupt {
			var score int

			pops := len(s)
			for i := 0; i < pops; i++ {
				var expected string
				expected, s = s.Pop()

				score = score * 5
				switch expected {
				case ")":
					score += 1
				case "]":
					score += 2
				case "}":
					score += 3
				case ">":
					score += 4
				}
			}

			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])
}
