package main

import "fmt"

func isValid(s string) bool {
	stack := make([]rune, 0)
	mapping := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if val, ok := mapping[char]; ok {
			topElement := ' '
			if len(stack) > 0 {
				topElement, stack = stack[len(stack)-1], stack[:len(stack)-1]
			}
			if val != topElement {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("{[]}"))   // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[}"))    // false
}
