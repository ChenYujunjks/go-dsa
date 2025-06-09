package stack

import (
	dll "github.com/ChenYujunjks/go-review/ds/link"
)

type Stack struct {
	list dll.LinkedList
}

func (s *Stack) Push(value any) {
	s.list.Append(value)
}
