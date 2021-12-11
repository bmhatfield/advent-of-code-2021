package main

import (
	"advent/aoc"
	"fmt"
	"strconv"
)

type point struct {
	x, y    int
	val     int
	flashed bool
}

func (p point) energized() bool {
	return p.val > 9
}

type grid struct {
	data [][]*point

	width  int
	height int
}

func (h grid) get(x, y int) *point {
	if x >= 0 && x < h.width && y >= 0 && y < h.height {
		return h.data[y][x]
	}

	return nil
}

func (h grid) adjacent(x, y int) []*point {
	points := make([]*point, 0, 8)

	coords := [][2]int{
		{x - 1, y},     // left
		{x + 1, y},     // right
		{x, y - 1},     // below
		{x, y + 1},     // above
		{x - 1, y - 1}, // b left
		{x - 1, y + 1}, // b right
		{x + 1, y - 1}, // a left
		{x + 1, y + 1}, // a right
	}

	for _, coord := range coords {
		p := h.get(coord[0], coord[1])
		if p != nil {
			points = append(points, p)
		}
	}

	return points
}

func (h grid) incr() []*point {
	full := make([]*point, 0)

	for _, pts := range h.data {
		for _, pt := range pts {
			pt.val++

			if pt.energized() {
				full = append(full, pt)
			}
		}
	}

	return full
}

func (h grid) reset() (flashed int) {
	for _, pts := range h.data {
		for _, pt := range pts {
			if pt.flashed {
				pt.val = 0
				pt.flashed = false
				flashed++
			}
		}
	}

	return flashed
}

func (h grid) String() string {
	var out string
	for _, rows := range h.data {
		for _, v := range rows {
			out += strconv.Itoa(v.val)
		}

		out += "\n"
	}

	return out
}

func newGrid(matrix [][]int) grid {
	g := grid{
		data:   make([][]*point, len(matrix)),
		width:  len(matrix[0]),
		height: len(matrix),
	}

	for y, vals := range matrix {
		g.data[y] = make([]*point, len(vals))

		for x, val := range vals {
			g.data[y][x] = &point{x: x, y: y, val: val}
		}
	}

	return g
}

func main() {
	matrix := aoc.Matrix(aoc.Lines("input"))

	g := newGrid(matrix)

	var syncd int
	for step := 0; step < 10000; step++ {
		fmt.Println(g)

		full := g.incr()
		if len(full) == 0 {
			continue
		}

		queue := make(chan *point, 1000)

		for _, pt := range full {
			queue <- pt
		}

		for pt := range queue {
			if pt.energized() && !pt.flashed {
				pt.flashed = true

				for _, adjPt := range g.adjacent(pt.x, pt.y) {
					adjPt.val++

					if adjPt.energized() && !adjPt.flashed {
						queue <- adjPt
					}
				}
			}

			if len(queue) == 0 {
				close(queue)
			}
		}

		flashed := g.reset()
		if flashed == 100 {
			syncd = step
			break
		}
	}

	fmt.Println(syncd)
}
