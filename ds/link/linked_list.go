package linkedlist

import "github.com/ChenYujunjks/go-review/ds/iface"

type Node struct {
	Value interface{}
	Next  *Node
}

// Node 方法实现了 NodeInter 接口
func (n *Node) GetValue() interface{} {
	return n.Value
}

func (n *Node) GetNext() iface.NodeInter {
	if n.Next == nil {
		return nil
	}
	return n.Next
}

type LinkedList struct {
	Head *Node
}

// LinkedList 方法实现了 LL 接口
func (l *LinkedList) GetHead() iface.NodeInter {
	return l.Head
}

// Go 中，你可以在定义结构体时只指定你想赋值的字段，其他字段会被自动初始化为该类型的零值
