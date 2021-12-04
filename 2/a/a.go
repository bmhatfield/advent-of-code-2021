package main

import (
	"advent/aoc"
	"fmt"
)

func main() {
	commands := aoc.Pairs(aoc.Lines("2.input"))

	var pos, depth int

	for _, command := range commands {
		switch command.Key {
		case "up":
			depth -= command.Value
		case "down":
			depth += command.Value
		case "forward":
			pos += command.Value
		default:
			panic(fmt.Errorf("unknown command: %s - %d", command.Key, command.Value))
		}
	}

	fmt.Printf("pos: %d, depth: %d, composite: %d", pos, depth, pos*depth)
}
