package main

import (
	"fmt"

	linkedlist "github.com/ChenYujunjks/go-review/ds/link"
	"github.com/ChenYujunjks/go-review/ds/list"
	"github.com/ChenYujunjks/go-review/ds/utils"
)

func main() {
	l := &linkedlist.LinkedList{Head: &linkedlist.Node{
		Value: nil,
		Next: &linkedlist.Node{
			Value: 42,
			Next:  nil,
		},
	}}
	l2 := &linkedlist.LinkedList{Head: &linkedlist.Node{
		Value: 42,
		Next:  nil,
	},
	}
	fmt.Println(utils.ToSlice(l))
	fmt.Println("LinkedList Length:", utils.Length(l))
	fmt.Println("------------------------------------------------")
	fmt.Println(utils.ToSlice(l2))
	fmt.Println("LinkedList Length:", utils.Length(l2))

	fmt.Println("List Part--------------------------------------------")

	l5 := list.NewList()
	l5.Append(1)
	l5.Append(2)
	l5.Append(3)
	l5.Prepend(10)
	fmt.Println("List Elements:", utils.ToSlice(l5))
	val, ok := l5.Get(3)
	if ok {
		fmt.Println("Get index 1:", val)
	} else {
		fmt.Println("Get index 1: not found")
	}
}
