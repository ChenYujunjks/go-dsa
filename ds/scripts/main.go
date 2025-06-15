package main

import (
	"fmt"

	linkedlist "github.com/ChenYujunjks/go-review/ds/link"
)

func main() {
	//l := list.NewList()
	//fmt.Println(l.GetHead() != nil)
	//result := utils.ToSlice(l)
	//fmt.Println(result)
	l1 := linkedlist.LinkedList{}
	fmt.Println(l1.GetHead())
	fmt.Println(l1.GetHead() == nil)
}
