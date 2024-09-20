package queue

import "fmt"

type Queue struct {
	Items []int
}

func (q *Queue) Enqueue(item int) {
	q.Items = append(q.Items, item)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty!")
	}
	deQueuedItem := q.Items[0]
	q.Items = q.Items[1:]
	return deQueuedItem
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty!")
		return -1
	}
	peekItem := q.Items[0]
	return peekItem
}

func (q *Queue) Size() int {
	return len(q.Items)
}
