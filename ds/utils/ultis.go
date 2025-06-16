package utils

import (
	"fmt"
	"reflect"

	"github.com/ChenYujunjks/go-review/ds/iface"
)

func PrintForward(l iface.LL) {
	current := l.GetHead()
	for current != nil {
		fmt.Printf("%d <-> ", current.GetValue())
		current = current.GetNext()
	}
	fmt.Println("nil")
}

func IsNilNode(n iface.NodeInter) bool {
	if n == nil {
		return true
	}
	return reflect.ValueOf(n).IsNil()
}


func Length(l iface.LL) int {
	count := 0
	current := l.GetHead()
	for !IsNilNode(current) {
		count++
		current = current.GetNext()
	}
	return count
}

// ToSlice 将链表转换为 any的切片
func ToSlice(l iface.LL) []any {
	var result []any
	current := l.GetHead()
	for !IsNilNode(current) {
		result = append(result, current.GetValue())
		current = current.GetNext()
	}
	return result
}
