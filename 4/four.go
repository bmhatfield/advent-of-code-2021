package four

import "advent/aoc"

func winner(s [5]int) bool {
	for _, v := range s {
		if v == 5 {
			return true
		}
	}

	return false
}

type Pos struct {
	marked bool
	x, y   int
}

type Board struct {
	numbers map[string]Pos

	rows [5]int
	cols [5]int
}

func (b *Board) Add(num string, x, y int) {
	b.numbers[num] = Pos{
		x: x,
		y: y,
	}
}

func (b *Board) Mark(num string) {
	pos, ok := b.numbers[num]
	if !ok {
		return
	}

	pos.marked = true

	// track marked columns/rows,
	// which allows much more efficient
	// checking of win state
	b.cols[pos.x]++
	b.rows[pos.y]++

	b.numbers[num] = pos
}

func (b Board) IsWin() bool {
	return winner(b.cols) || winner(b.rows)
}

func (b Board) SumUnmarked() (out int) {
	for num, pos := range b.numbers {
		if pos.marked {
			continue
		}

		out += aoc.Int(num)
	}

	return out
}

func NewBoard() *Board {
	return &Board{
		numbers: make(map[string]Pos),
	}
}
