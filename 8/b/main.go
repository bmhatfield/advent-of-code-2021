package main

import (
	"advent/aoc"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Char string

func (s Char) Segments() []string {
	return strings.Split(string(s), "")
}

func (s Char) InclAllOf(o Char) bool {
	for _, l := range o.Segments() {
		if !strings.Contains(string(s), l) {
			return false
		}
	}

	return true
}

func (s Char) Excl(o Char) Char {
	var e string
	for _, l := range o.Segments() {
		if !strings.Contains(string(s), l) {
			e += l
		}
	}

	return AsChar(e)
}

func AsChar(s string) Char {
	cs := strings.Split(s, "")
	sort.Strings(cs)
	return Char(strings.Join(cs, ""))
}

func signals(chars []string) map[Char]int {
	mapping := make(map[Char]int)
	c := make(map[int]Char)

	unclear := [6]Char{}
	var uc int
	for _, raw := range chars {
		char := AsChar(raw)

		switch len(raw) {
		case 2:
			mapping[char] = 1
			c[1] = char
		case 4:
			mapping[char] = 4
			c[4] = char
		case 3:
			mapping[char] = 7
			c[7] = char
		case 7:
			mapping[char] = 8
			c[9] = char
		default:
			unclear[uc] = char
			uc++
		}
	}

	ft := [2]Char{}
	var ftp int
	for _, char := range unclear {
		switch len(char) {
		case 6: // 0, 9, 6
			switch {
			case !char.InclAllOf(c[4]) && char.InclAllOf(c[1]): // 0
				mapping[char] = 0
				c[0] = char
			case !char.InclAllOf(c[1]): // 6
				mapping[char] = 6
				c[6] = char
			default: // 9
				mapping[char] = 9
				c[9] = char
			}
		case 5: // 5, 2, 3
			switch {
			case char.InclAllOf(c[1]): // 3
				mapping[char] = 3
				c[3] = char
			default:
				ft[ftp] = char
				ftp++
			}
		}
	}

	for _, char := range ft {
		switch {
		case c[6].InclAllOf(char):
			mapping[char] = 5
			c[5] = char
		default:
			mapping[char] = 2
			c[2] = char
		}
	}

	return mapping
}

func output(chars []string, mapping map[Char]int) int {
	var v int
	for i, char := range chars {
		n := mapping[AsChar(char)]
		v += n * int(math.Pow10(len(chars)-1-i))
	}

	fmt.Println(chars, mapping, v)
	return v
}

func main() {
	scrambled := aoc.SignalOutputs(aoc.Lines("input"))

	var sum int
	for _, scram := range scrambled {
		chars := signals(scram[0])
		num := output(scram[1], chars)

		sum += num
	}

	fmt.Println(sum)
}
