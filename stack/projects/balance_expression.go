package projects

import "stack/stack"

func IsBalanced(expression string) bool {
	st := stack.Stack{}

	mapItems := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range expression {
		if isExistsInMap(char, mapItems) {
			st.Push(char)
		}

		if _, exists := mapItems[char]; exists {
			pop, err := st.Pop()
			if err != nil {
				return false
			}
			if pop != mapItems[char] {
				return false
			}
		}
	}

	return st.IsEmpty()
}

func isExistsInMap(char rune, m map[rune]rune) bool {
	exists := false
	for _, value := range m {
		if value == char {
			exists = true
			break
		}
	}
	return exists
}
