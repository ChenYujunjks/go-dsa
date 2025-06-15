package linkedlist

import "fmt"



func PrintForward(l LL) {
	current := l.GetHead()
	for current != nil {
		fmt.Printf("%d <-> ", current.GetValue())
		current = current.GetNext()
	}
	fmt.Println("nil")
}

// Length 返回链表的长度
func Length(l LL) int {
	count := 0
	current := l.GetHead()
	for current != nil {
		count++
		current = current.GetNext()
	}
	return count
}

// ToSlice 将链表转换为 any的切片
func ToSlice(l LL) []any {
	var result []any
	current := l.GetHead()
	for current != nil {
		result = append(result, current.GetValue())
		current = current.GetNext()
	}
	return result
}
