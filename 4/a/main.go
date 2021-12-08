package main

import (
	four "advent/4"
	"advent/aoc"
	"fmt"
)

func main() {
	drawn, brds := aoc.Boards(aoc.Lines("4.input"))

	boards := []*four.Board{}

	for _, brd := range brds {
		b := four.NewBoard()

		for y, row := range brd {
			for x, col := range row {
				b.Add(col, x, y)
			}
		}

		boards = append(boards, b)
	}

	for _, num := range drawn {
		for _, b := range boards {
			b.Mark(num)

			if b.IsWin() {
				sum := b.SumUnmarked()
				fmt.Println(sum, num, sum*aoc.Int(num))
				return
			}
		}
	}

	fmt.Println(boards)
}
