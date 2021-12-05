package five

import (
	"advent/aoc"

	"github.com/pkg/math"
)

type Plot map[Point]int

func (p Plot) AddLine(line []Point) {
	for _, point := range line {
		p[point]++
	}
}

func (p Plot) Intersects() (out int) {
	for _, num := range p {
		if num > 1 {
			out++
		}
	}

	return out
}

type Point struct {
	X int
	Y int
}

func (p Point) Distance(o Point) int {
	return math.MaxInt(
		aoc.Abs(p.X-o.X),
		aoc.Abs(p.Y-o.Y),
	)
}

func (p Point) nv(a, b int) int {
	switch {
	case a > b:
		return a - 1
	case a < b:
		return a + 1
	default:
		return a
	}
}

func (p Point) Next(end Point) Point {
	p.X = p.nv(p.X, end.X)
	p.Y = p.nv(p.Y, end.Y)

	return p
}

type Vector struct {
	Start Point
	End   Point
}

func (v Vector) length() int {
	return v.Start.Distance(v.End)
}

func (v Vector) horizontal() bool {
	return v.Start.Y == v.End.Y
}

func (v Vector) vertical() bool {
	return v.Start.X == v.End.X
}

func (v Vector) IsStraight() bool {
	return v.horizontal() || v.vertical()
}

func (v Vector) Line() []Point {
	length := v.length()
	pts := make([]Point, length+1)

	point := v.Start
	for i := 0; i <= length; i++ {
		pts[i] = point
		point = point.Next(v.End)
	}

	return pts
}

func NewVector(raw [2][2]int) Vector {
	return Vector{
		Start: Point{
			X: raw[0][0],
			Y: raw[0][1],
		},
		End: Point{
			X: raw[1][0],
			Y: raw[1][1],
		},
	}
}
