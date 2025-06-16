package stack

import (
	"github.com/ChenYujunjks/go-review/ds/iface"
	l "github.com/ChenYujunjks/go-review/ds/link"
)

type Stack struct {
	list *l.LinkedList
	len  int
}

func New() *Stack {
	// 创建一个空栈
	return &Stack{
		list: &l.LinkedList{},
		len:  0,
	}
}

// Push 将元素压入栈顶（链表头部）
func (s *Stack) Push(val interface{}) {
	newNode := &l.Node{Value: val, Next: s.list.Head}
	s.list.Head = newNode
	s.len++
}

func (s *Stack) Pop() (any, bool) {
	if s.list.Head == nil {
		return nil, false // 栈为空
	}

	val := s.list.Head.Value
	s.list.Head = s.list.Head.Next
	s.len--

	return val, true
}

func (s *Stack) Peek() (any, bool) {
	if s.list.Head == nil {
		return nil, false // 栈为空
	}

	return s.list.Head.Value, true
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}
func (s *Stack) Len() int {
	return s.len
}
func (s *Stack) GetHead() iface.NodeInter {
	return s.list.Head
}
