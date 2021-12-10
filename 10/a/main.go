package main

import (
	ten "advent/10"
	"advent/aoc"
	"fmt"
	"strings"
)

func main() {
	lines := aoc.Lines("input")

	var score int

	for _, line := range lines {
		s := ten.Stack{}
		fmt.Println(line)

		var done bool
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
					fmt.Println(ch, expected)

					switch ch {
					case ")":
						score += 3
					case "]":
						score += 57
					case "}":
						score += 1197
					case ">":
						score += 25137
					}

					done = true
				}
			}

			if done {
				break
			}
		}
	}

	fmt.Println(score)
}
