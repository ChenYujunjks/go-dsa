package queue

import (
	"fmt"

	dll "github.com/ChenYujunjks/go-review/ds/link"
)

type Queue struct {
	list dll.LinkedList
}

// Enqueue 添加元素到队尾
func (q *Queue) Enqueue(value any) {
	q.list.Append(value)
}

// Dequeue 移除队首元素并返回
func (q *Queue) Dequeue() any {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}
	value := q.list.Head.Value
	q.list.Delete(value) // 删除的是头节点的值
	return value
}

// Peek 查看队首元素但不移除
func (q *Queue) Peek() any {
	if q.IsEmpty() {
		return nil
	}
	return q.list.Head.Value
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.list.Head == nil
}

// Length 返回队列长度
func (q *Queue) Length() int {
	return dll.Length(q.list)
}

// ToSlice 返回队列所有元素
func (q *Queue) ToSlice() []any {
	return q.list.ToSlice()
}

func (q *Queue) Print() {
	current := q.list.Head
	for current != nil {
		fmt.Printf("%v <- ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}
