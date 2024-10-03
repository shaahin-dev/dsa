package hashmap

import "fmt"

type Node struct {
	key   string
	value int
	next  *Node
}

type hashMap struct {
	buckets []*Node
	size    int
}

func NewHashMap(size int) *hashMap {
	return &hashMap{
		buckets: make([]*Node, size),
		size:    size,
	}
}

func (hm *hashMap) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % hm.size
}

func (hm *hashMap) Insert(key string, value int) {
	index := hm.hash(key)
	node := hm.buckets[index]

	if node == nil {
		hm.buckets[index] = &Node{
			key:   key,
			value: value,
		}
		return
	}

	for node != nil {
		if node.key == key {
			node.value = value
			return
		}
		if node.next == nil {
			break
		}
		node = node.next
	}

	node.next = &Node{
		key:   key,
		value: value,
	}
}

func (hm *hashMap) Get(key string) (int, bool) {
	index := hm.hash(key)
	node := hm.buckets[index]

	for node != nil {
		if node.key == key {
			return node.value, true
		}
		node = node.next
	}
	return 0, false
}

func (hm *hashMap) Delete(key string) {
	index := hm.hash(key)
	node := hm.buckets[index]

	if node == nil {
		return
	}

	if node.key == key {
		hm.buckets[index] = node.next
		return
	}

	prev := node
	for node != nil {
		if node.key == key {
			prev.next = node.next
			return
		}
		prev = node
		node = node.next
	}

}

func (hm *hashMap) Display() {
	for i, node := range hm.buckets {
		fmt.Printf("Bucket %d: ", i)
		for node != nil {
			fmt.Printf("(%s: %d) -> ", node.key, node.value)
			node = node.next
		}
		fmt.Println()
	}
}
