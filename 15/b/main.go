package main

import (
	f "advent/15"
	"advent/aoc"
	"fmt"
)

func score(pts []*f.Point) uint {
	var s uint

	for _, pt := range pts {
		s += uint(pt.Val)
	}

	return s
}

func path(grid f.Grid) uint {
	start := grid.Get(0, 0)

	prev := make(map[f.Point]f.Point)

	cost := map[f.Point]uint{
		*start: 0,
	}

	q := make(chan f.Point, 100000)
	q <- *start

	for cur := range q {
		for _, n := range grid.Adjacent(cur.X, cur.Y) {
			neighbor := *n

			nCost := cost[cur] + uint(neighbor.Val)

			bestCost, ok := cost[neighbor]
			if !ok || nCost < bestCost {
				prev[neighbor] = cur
				cost[neighbor] = nCost

				q <- neighbor
			}
		}

		if len(q) == 0 {
			close(q)
		}
	}

	return cost[*grid.Get(grid.Width-1, grid.Height-1)]
}

func main() {
	trix := aoc.Matrix(aoc.Lines("input"))

	newtrix := make([][]int, len(trix)*5)
	for y, row := range trix {
		for yo := 0; yo < 5; yo++ {
			newtrix[y+(yo*len(trix))] = make([]int, len(row)*5)

			for x, val := range row {
				for xo := 0; xo < 5; xo++ {
					nv := (val + yo + xo)
					if nv > 9 {
						nv = nv % 9
					}

					newtrix[y+(yo*len(trix))][x+(xo*len(row))] = nv
				}
			}
		}
	}

	grid := f.NewGrid(newtrix)

	fmt.Println(path(grid))
}
