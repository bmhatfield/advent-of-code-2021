package main

import (
	four "advent/4"
	"advent/aoc"
	"fmt"
)

func isLast(bs []bool) bool {
	var seen bool

	for _, b := range bs {
		if !b {
			if seen {
				return false
			}

			seen = true
		}
	}

	return true
}

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

	winners := make([]bool, len(boards))

	for _, num := range drawn {
		for i, b := range boards {
			if winners[i] {
				continue
			}

			b.Mark(num)

			if b.IsWin() {
				if isLast(winners) {
					sum := b.SumUnmarked()
					fmt.Println(sum, num, sum*aoc.Int(num))
					return
				}

				winners[i] = true
			}
		}
	}

	fmt.Println(boards)
}
