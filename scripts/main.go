package main

import (
	dll "github.com/ChenYujunjks/go-review/ds/link"
	"github.com/ChenYujunjks/go-review/ds/queue"
)

func test_link() {
	list := dll.LinkedList{}

	list.Append(10)
	list.Append(20)
	list.Prepend(5)

	list.PrintForward() // 5 <-> 10 <-> 20 <-> nil

	list.Delete(10)
	list.PrintForward() // 5 <-> 20 <-> nil
}

func test_queue() {
	queue := queue.Queue{}

	queue.Enqueue(1)
	queue.Enqueue("apple")
	queue.Enqueue(3)

	queue.Print() // 1 -> 2 -> 3 -> nil

	dequeued := queue.Dequeue()
	println("Dequeued:", dequeued) // Dequeued: 1

	queue.Print() // 2 -> 3 -> nil
}
func main() {
	//test_link()
	//stest_queue()
}
