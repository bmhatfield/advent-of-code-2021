package twelve

import "fmt"

const (
	Start = "start"
	End   = "end"
)

type Index map[string]map[string]struct{}

func (i Index) ins(l, r string) {
	if i[l] == nil {
		i[l] = make(map[string]struct{})
	}

	i[l][r] = struct{}{}
}

func (i Index) Insert(l, r string) {
	if l != End && r != Start {
		i.ins(l, r)
	}

	if r != End && l != Start {
		i.ins(r, l)
	}
}

func (i Index) String() string {
	var s string

	for l, rights := range i {
		s += fmt.Sprintln(l)

		for r := range rights {
			s += fmt.Sprintln("-", r)
		}
	}

	return s
}

func NewIndex() Index {
	return make(Index)
}
