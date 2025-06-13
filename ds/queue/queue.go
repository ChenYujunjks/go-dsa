package queue

import (
	l "github.com/ChenYujunjks/go-review/ds/link"
)

type Queue struct {
	list *l.LinkedList
	tail *l.Node // 为了快速 enqueue
	len  int
}

// New 创建一个空队列 **封装细节**
func New() *Queue {
	return &Queue{
		list: &l.LinkedList{},
		tail: nil,
		len:  0,
	}
}

func (q *Queue) Enqueue(value interface{}) {
	newNode := &l.Node{Value: value}
	if q.list.Head == nil {
		q.list.Head = newNode
		q.tail = newNode
	} else {
		q.tail.Next = newNode
		q.tail = newNode
	}
	q.len++
}

// Dequeue 移除队首元素并返回
func (q *Queue) Dequeue() (interface{}, bool) {
	if q.list.Head == nil {
		return nil, false
	}

	val := q.list.Head.Value
	q.list.Head = q.list.Head.Next
	q.len--

	if q.list.Head == nil {
		q.tail = nil // 队列清空，tail 也置空
	}
	return val, true
}

func (q *Queue) Peek() (interface{}, bool) {
	if q.list.Head == nil {
		return nil, false
	}
	return q.list.Head.Value, true
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

func (q *Queue) Len() int {
	return q.len
}
