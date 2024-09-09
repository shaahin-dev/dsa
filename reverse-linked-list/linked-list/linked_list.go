package linked_list

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) Show() {
	current := ll.head
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) Add(value interface{}) {
	newNode := &Node{value: value}
	if ll.tail == nil {
		ll.tail = newNode
		ll.head = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}
}

func (ll *LinkedList) Remove(key interface{}) {
	if ll.head == nil {
		return
	}
	for ll.head != nil && ll.head.value == key {
		ll.head = ll.head.next
	}
	if ll.head == nil {
		ll.tail = nil
		return
	}
	current := ll.head
	for current != nil && current.next != nil {
		if current.next.value == key {
			current.next = current.next.next
			if current.next == nil {
				ll.tail = current
			}
		} else {
			current = current.next
		}
	}
}

func (ll *LinkedList) Find(vale interface{}) *Node {
	current := ll.head
	for current != nil {
		if current.value == vale {
			return current
		}
		current = current.next
	}
	return nil
}

func (ll *LinkedList) Reverse() {
	var prev, next *Node
	current := ll.head
	ll.tail = ll.head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	ll.head = prev
}
