package thirteen

import (
	"advent/aoc"
	"strings"
)

func Plot(grid [][]bool, coords [][2]int) [][]bool {
	for _, coord := range coords {
		x, y := coord[0], coord[1]

		grid[y][x] = true
	}

	return grid
}

func MkGrid(height, width int) [][]bool {
	grid := make([][]bool, height)
	for y := range grid {
		grid[y] = make([]bool, width)
	}
	return grid
}

func Transform(x, y, fx, fy int) (int, int) {
	if fx > 0 {
		x = fx - (x - fx)
	}

	if fy > 0 {
		y = fy - (x - fy)
	}

	return x, y
}

func FoldY(grid [][]bool, fy int) [][]bool {
	out := MkGrid(fy, len(grid[0]))

	for y, row := range grid {
		for x, col := range row {
			if y > fy {
				ny := fy - (y - fy)

				out[ny][x] = out[ny][x] || col

				continue
			}

			if y == fy {
				continue
			}

			out[y][x] = col
		}
	}

	return out
}

func FoldX(grid [][]bool, fx int) [][]bool {
	out := MkGrid(len(grid), fx)

	for y, row := range grid {
		for x, col := range row {
			if x > fx {
				nx := fx - (x - fx)

				out[y][nx] = out[y][nx] || col

				continue
			}

			if x == fx {
				continue
			}

			out[y][x] = col
		}
	}

	return out
}

func Dimensions(coords [][2]int) (int, int) {
	var h, w int
	for _, coord := range coords {
		if coord[0] > w {
			w = coord[0]
		}

		if coord[1] > h {
			h = coord[1]
		}
	}
	return h + 1, w + 1
}

func SGrid(grid [][]bool) string {
	var s string
	for _, row := range grid {
		for _, col := range row {
			if col {
				s += "#"
			} else {
				s += "."
			}
		}

		s += "\n"
	}
	return s
}

func Fold(grid [][]bool, instruction string) [][]bool {
	line := strings.Split(instruction, " ")
	equa := strings.Split(line[2], "=")
	axis, val := equa[0], aoc.Int(equa[1])

	switch axis {
	case "x":
		return FoldX(grid, val)
	case "y":
		return FoldY(grid, val)
	default:
		panic("invalid")
	}
}
