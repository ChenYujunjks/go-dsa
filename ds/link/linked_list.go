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
