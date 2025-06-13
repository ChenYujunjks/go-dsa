package linkedlist

import "fmt"

func (l *LinkedList) PrintForward() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d <-> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

// Length 返回链表的长度
func (l *LinkedList) Length() int {
	count := 0
	current := l.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

// ToSlice 将链表转换为 any的切片
func (l *LinkedList) ToSlice() []any {
	var result []any
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}
