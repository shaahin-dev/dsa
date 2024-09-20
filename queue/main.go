package main

import (
	"fmt"
	"queue/queue"
)

func main() {
	q := queue.Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("Queue after enqueue:", q.Items)

	fmt.Println("Dequeued:", q.Dequeue())
	fmt.Println("Queue after dequeue:", q.Items)

	fmt.Println("Peek:", q.Peek())

	fmt.Println("Is empty:", q.IsEmpty())

	fmt.Println("Queue size:", q.Size())
}
