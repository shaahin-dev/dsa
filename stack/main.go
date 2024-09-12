package main

import (
	"fmt"
	"stack/stack"
)

func main() {
	stack := stack.Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	peek, err := stack.Peek()
	if err != nil {
		return
	}
	fmt.Println(peek)
	pop, err := stack.Pop()
	if err != nil {
		return
	}
	fmt.Println(pop)
	stack.Show()
}
