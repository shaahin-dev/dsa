package queue

import "fmt"

type Item struct {
	Value    string
	Priority int
}

type PriorityQueue struct {
	Items []Item
}

func (pq *PriorityQueue) Enqueue(value string, priority int) {
	item := Item{Value: value, Priority: priority}

	index := 0
	for index < len(pq.Items) && pq.Items[index].Priority < priority {
		index++
	}

	pq.Items = append(pq.Items[:index], append([]Item{item}, pq.Items[index:]...)...)
}

func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.Items) == 0
}

func (pq *PriorityQueue) Dequeue() *Item {
	if pq.IsEmpty() {
		fmt.Println("Queue is empty")
		return nil
	}
	highestPriorityItem := pq.Items[0]
	pq.Items = pq.Items[1:]
	return &highestPriorityItem
}
func (pq *PriorityQueue) Peek() *Item {
	if len(pq.Items) == 0 {
		fmt.Println("Queue is empty")
		return nil
	}

	return &pq.Items[0]
}

func (pq *PriorityQueue) Size() int {
	return len(pq.Items)
}
