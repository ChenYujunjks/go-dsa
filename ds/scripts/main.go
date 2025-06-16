package main

import (
	"fmt"
	"log"

	linkedlist "github.com/ChenYujunjks/go-review/ds/link"
	"github.com/ChenYujunjks/go-review/ds/stack"
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

	fmt.Println("Stack LL Part--------------------------------------------")

	s := stack.New()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	slice := utils.ToSlice(s)

	expected := []any{30, 20, 10}
	for i, val := range expected {
		if slice[i] != val {
			log.Fatalf("❌ Expected %v at index %d, got %v", val, i, slice[i])
		}
	}

	fmt.Println("✅ All elements matched expected stack order.")
}
