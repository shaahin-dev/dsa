package stack

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	topIndex := len(s.items) - 1
	topItem := s.items[topIndex]
	s.items = s.items[:topIndex]
	return topItem, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	topIndex := len(s.items) - 1
	return s.items[topIndex], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Show() {
	for _, v := range s.items {
		println(v)
	}
}
