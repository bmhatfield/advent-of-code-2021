package fifteen

import "strconv"

type Point struct {
	X, Y int
	Val  int
}

type Grid struct {
	Points [][]*Point

	Width  int
	Height int
}

func (h Grid) Get(x, y int) *Point {
	if x >= 0 && x < h.Width && y >= 0 && y < h.Height {
		return h.Points[y][x]
	}

	return nil
}

func (h Grid) Adjacent(x, y int) []*Point {
	points := make([]*Point, 0, 4)

	coords := [][2]int{
		{x - 1, y}, // left
		{x + 1, y}, // right
		{x, y - 1}, // below
		{x, y + 1}, // above
	}

	for _, coord := range coords {
		p := h.Get(coord[0], coord[1])
		if p != nil {
			points = append(points, p)
		}
	}

	return points
}

func (h Grid) String() string {
	var out string
	for _, rows := range h.Points {
		for _, v := range rows {
			out += strconv.Itoa(v.Val)
		}

		out += "\n"
	}

	return out
}

func NewGrid(matrix [][]int) Grid {
	g := Grid{
		Points: make([][]*Point, len(matrix)),
		Width:  len(matrix[0]),
		Height: len(matrix),
	}

	for y, vals := range matrix {
		g.Points[y] = make([]*Point, len(vals))

		for x, val := range vals {
			g.Points[y][x] = &Point{X: x, Y: y, Val: val}
		}
	}

	return g
}
