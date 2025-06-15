package stack

import (
	l "github.com/ChenYujunjks/go-review/ds/link"
)

type Stack struct {
	list l.LinkedList
	len  int
}

// Push 将元素压入栈顶（链表头部）
func (s *Stack) Push(val interface{}) {
	newNode := &l.Node{Value: val, Next: s.list.Head}
	s.list.Head = newNode
	s.len++
}
