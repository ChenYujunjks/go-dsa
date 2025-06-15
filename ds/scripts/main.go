package main

import (
	"fmt"

	linkedlist "github.com/ChenYujunjks/go-review/ds/link"
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
}
