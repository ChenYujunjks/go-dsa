package linkedlist

type Node struct {
	Value interface{}
	Next  *Node
}

func (n *Node) GetValue() interface{} {
	return n.Value
}

func (n *Node) GetNext() NodeInter {
	if n.Next == nil {
		return nil
	}
	return n.Next
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) GetHead() NodeInter {
	return l.Head
}

// Go 中，你可以在定义结构体时只指定你想赋值的字段，其他字段会被自动初始化为该类型的零值
// Append adds a node at the end of the list.
func (l *LinkedList) Append(value interface{}) {
	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Prepend adds a node at the beginning of the list.
func (l *LinkedList) Prepend(value interface{}) {
	newNode := &Node{Value: value, Next: l.Head}
	l.Head = newNode
}

// Find returns the first node with the matching value.
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

// Delete removes the first node with the given value.
func (l *LinkedList) Delete(value interface{}) bool {
	if l.Head == nil {
		return false
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return true
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}
	return false
}

// Insert inserts value at the given index (0-based).
// If index > current length, it appends to the end.
func (l *LinkedList) Insert(index int, value interface{}) {
	newNode := &Node{Value: value}

	if index <= 0 || l.Head == nil {
		// Prepend at head
		newNode.Next = l.Head
		l.Head = newNode
		return
	}

	current := l.Head
	pos := 0
	for current.Next != nil && pos < index-1 {
		current = current.Next
		pos++
	}

	newNode.Next = current.Next
	current.Next = newNode
}
