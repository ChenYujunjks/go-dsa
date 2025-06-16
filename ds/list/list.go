package list

import (
	"github.com/ChenYujunjks/go-review/ds/iface"
	l "github.com/ChenYujunjks/go-review/ds/link"
)

type List struct {
	ll *l.LinkedList
}

func NewList() *List {
	return &List{ll: &l.LinkedList{}}
}

func (lst *List) GetHead() iface.NodeInter {
	return lst.ll.Head
}

// Go 中，你可以在定义结构体时只指定你想赋值的字段，其他字段会被自动初始化为该类型的零值
// Append adds a node at the end of the list.
func (lst *List) Append(value interface{}) {
	newNode := &l.Node{Value: value}
	if lst.ll.Head == nil {
		lst.ll.Head = newNode
		return
	}
	current := lst.ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (lst *List) Prepend(value interface{}) {
	newNode := &l.Node{Value: value, Next: lst.ll.Head}
	lst.ll.Head = newNode
}

func (lst *List) Find(value interface{}) *l.Node {
	current := lst.ll.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

// Delete removes the first node with the given value.
func (lst *List) Delete(value interface{}) bool {
	if lst.ll.Head == nil {
		return false
	}
	if lst.ll.Head.Value == value {
		lst.ll.Head = lst.ll.Head.Next
		return true
	}
	current := lst.ll.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}
	return false
}

func (lst *List) Get(index int) (interface{}, bool) {
	if index < 0 {
		return nil, false
	}
	current := lst.ll.Head
	for i := 0; current != nil && i < index; i++ {
		current = current.Next
	}
	if current == nil {
		return nil, false
	}
	return current.Value, true
}

// Insert inserts value at the given index (0-based).
// If index > current length, it appends to the end.
func (lst *List) Insert(index int, value interface{}) bool {
	if index < 0 {
		return false
	}
	if index == 0 {
		lst.Prepend(value)
		return true
	}
	current := lst.ll.Head
	for i := 0; current != nil && i < index-1; i++ {
		current = current.Next
	}
	if current == nil {
		return false
	}
	newNode := &l.Node{Value: value, Next: current.Next}
	current.Next = newNode
	return true
}
