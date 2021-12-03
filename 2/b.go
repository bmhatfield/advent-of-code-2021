package main

import (
	"advent/aoc"
	"fmt"
)

func main() {
	commands := aoc.Pairs(aoc.Lines("2.input"))

	var pos, depth, aim int

	for _, command := range commands {
		switch command.Key {
		case "up":
			aim -= command.Value
		case "down":
			aim += command.Value
		case "forward":
			pos += command.Value
			depth += aim * command.Value
		default:
			panic(fmt.Errorf("unknown command: %s - %d", command.Key, command.Value))
		}
	}

	fmt.Printf("pos: %d, depth: %d, composite: %d", pos, depth, pos*depth)
}
