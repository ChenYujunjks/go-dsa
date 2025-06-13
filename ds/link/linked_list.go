package linkedlist

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

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
