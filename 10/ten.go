package ten

type Stack []string

func (s Stack) Push(v string) Stack {
	s = append(s, v)
	return s
}

func (s Stack) Pop() (string, Stack) {
	n := len(s) - 1
	return s[n], s[:n]
}
