package main

import (
	"advent/aoc"
	"fmt"
	"sort"
)

type point struct {
	x, y, height int
}

type heightmap struct {
	data [][]int

	width  int
	height int
}

func (h heightmap) get(x, y int) point {
	return point{
		x:      x,
		y:      y,
		height: h.data[y][x],
	}
}

func (h heightmap) walk(x, y int) []point {
	points := make([]point, 0, 4)

	if x > 0 {
		points = append(points, h.get(x-1, y))
	}

	if x < h.width-1 {
		points = append(points, h.get(x+1, y))
	}

	if y > 0 {
		points = append(points, h.get(x, y-1))
	}

	if y < h.height-1 {
		points = append(points, h.get(x, y+1))
	}

	return points
}

func (h heightmap) isLow(x, y int) bool {
	pt := h.get(x, y)

	for _, point := range h.walk(x, y) {
		if pt.height >= point.height {
			return false
		}
	}

	return true
}

func (h heightmap) basin(x, y int) map[point]struct{} {
	low := h.get(x, y)

	basin := make(map[point]struct{})
	q := make(chan point, 100)

	add := func(p point) {
		q <- p
		basin[p] = struct{}{}
	}

	// seed our first point, the lowpoint of the basin
	add(low)

	// iterate all point branches
	for point := range q {
		for _, next := range h.walk(point.x, point.y) {
			if next.height >= point.height && next.height != 9 {
				if _, ok := basin[next]; !ok {
					add(next)
				}
			}
		}

		if len(q) == 0 {
			close(q)
		}
	}

	return basin
}

func main() {
	h := heightmap{
		data: aoc.Matrix(aoc.Lines("input")),
	}
	h.width = len(h.data[0])
	h.height = len(h.data)

	lows := make([][2]int, 0)
	for x := 0; x < h.width; x++ {
		for y := 0; y < h.height; y++ {
			if h.isLow(x, y) {
				lows = append(lows, [2]int{x, y})
			}
		}
	}

	sizes := make([]int, len(lows))
	for b, low := range lows {
		sizes[b] = len(h.basin(low[0], low[1]))
	}

	sort.Ints(sizes)

	mult := 1
	for _, i := range sizes[len(sizes)-3:] {
		mult = mult * i
	}

	fmt.Println(mult)
}
