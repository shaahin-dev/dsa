package main

import (
	"fmt"
	"stack/projects"
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
	reverseString := projects.ReverseString("abcdef")
	fmt.Println(reverseString)
	fmt.Println(projects.IsBalanced("(2+2)*[(2*4)]"))
	fmt.Println(projects.IsBalanced("(2+2)*[(2*4]"))
}
