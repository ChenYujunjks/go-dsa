package linkedlist

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Append(value interface{}) { // 尾部插入
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	l.Tail.Next = newNode
	newNode.Prev = l.Tail
	l.Tail = newNode

}
func (l *LinkedList) Prepend(value interface{}) { // 头部插入
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode
}

func (l *LinkedList) Find(value interface{}) *Node {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *LinkedList) Delete(value interface{}) bool {
	current := l.Head
	for current != nil {
		if current.Value == value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				l.Head = current.Next // 如果是头节点
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				l.Tail = current.Prev // 如果是尾节点
			}
			return true
		}
		current = current.Next
	}
	return false
}
func (l *LinkedList) PrintForward() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d <-> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func (l *LinkedList) PrintBackward() {
	current := l.Tail
	for current != nil {
		fmt.Printf("%d <-> ", current.Value)
		current = current.Prev
	}
	fmt.Println("nil")
}
