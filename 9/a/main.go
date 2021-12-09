package main

import (
	"advent/aoc"
	"fmt"
)

type heightmap struct {
	data [][]int

	width  int
	height int
}

func (h heightmap) get(x, y int) int {
	return h.data[y][x]
}

func (h heightmap) isLow(x, y int) bool {
	pt := h.get(x, y)

	switch {
	case x > 0 && h.get(x-1, y) <= pt:
		return false
	case x+1 < h.width && h.get(x+1, y) <= pt:
		return false
	case y > 0 && h.get(x, y-1) <= pt:
		return false
	case y+1 < h.height && h.get(x, y+1) <= pt:
		return false
	}

	return true
}

func main() {
	h := heightmap{
		data: aoc.Matrix(aoc.Lines("input")),
	}
	h.width = len(h.data[0])
	h.height = len(h.data)

	lows := make([][2]int, 0)
	var sum int

	for x := 0; x < h.width; x++ {
		for y := 0; y < h.height; y++ {
			if h.isLow(x, y) {
				lows = append(lows, [2]int{x, y})
				sum += 1 + h.get(x, y)
			}
		}
	}

	fmt.Println(lows)
	fmt.Println(sum)
}
