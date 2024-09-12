package projects

import "stack/stack"

func ReverseString(str string) string {
	st := stack.Stack{}
	for _, ch := range str {
		st.Push(ch)
	}
	var reverseString []rune
	for !st.IsEmpty() {
		popChar, err := st.Pop()
		if err != nil {
			return "pop of stack error"
		}
		reverseString = append(reverseString, popChar)
	}
	return string(reverseString)
}
