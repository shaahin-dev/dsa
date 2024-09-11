package main

import linked_list "linked-list/linked-list"

func main() {
	ll := &linked_list.LinkedList{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.Remove(1)
	ll.Reverse()
	ll.Show()
	ll.FindKthFromEnd(2)
}
