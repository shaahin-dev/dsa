package main

import (
	"fmt"
	"queue/queue"
)

func main() {
	q := &queue.Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("Queue after enqueue:", q.Items)

	fmt.Println("Dequeued:", q.Dequeue())
	fmt.Println("Queue after dequeue:", q.Items)

	fmt.Println("Peek:", q.Peek())

	fmt.Println("Is empty:", q.IsEmpty())

	fmt.Println("Queue size:", q.Size())

	pq := &queue.PriorityQueue{}

	pq.Enqueue("Task 1", 3) // Priority 3 (low)
	pq.Enqueue("Task 2", 1) // Priority 1 (high)
	pq.Enqueue("Task 3", 2) // Priority 2 (medium)

	fmt.Println("Priority Queue:", pq.Items)

	peekedItem := pq.Peek()
	if peekedItem != nil {
		fmt.Println("Peek highest priority:", peekedItem.Value, "with priority", peekedItem.Priority)
	}

	for !pq.IsEmpty() {
		item := pq.Dequeue()
		if item != nil {
			fmt.Println("Dequeued:", item.Value, "with priority", item.Priority)
		}
	}
}
