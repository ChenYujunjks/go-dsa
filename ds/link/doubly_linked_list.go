package linkedlist

type DNode struct {
	Value interface{}
	Next  *DNode
	Prev  *DNode
}

type DoublyLinkedList struct {
	Head *DNode
	Tail *DNode
}

func (n *DNode) GetValue() interface{} {
	return n.Value
}

func (n *DNode) GetNext() NodeInter {
	if n.Next == nil {
		return nil
	}
	return n.Next
}

func (l *DoublyLinkedList) Append(value interface{}) {
	newNode := &DNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	l.Tail.Next = newNode
	newNode.Prev = l.Tail
	l.Tail = newNode
}

func (l *DoublyLinkedList) Prepend(value interface{}) {
	newNode := &DNode{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode
}

func (l *DoublyLinkedList) Find(value interface{}) *DNode {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *DoublyLinkedList) Delete(value interface{}) bool {
	current := l.Head
	for current != nil {
		if current.Value == value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				l.Head = current.Next // 删除头节点
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				l.Tail = current.Prev // 删除尾节点
			}
			return true
		}
		current = current.Next
	}
	return false
}
