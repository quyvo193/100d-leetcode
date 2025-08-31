package main

import "fmt"

// {([])}
func isValid(s string) bool {
	stack := make([]rune, len(s))

	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (c == '}' && pop != '{') ||
				(c == ']' && pop != '[') ||
				(c == ')' && pop != '(') {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func main() {
	r := isValid("()[]{}")
	fmt.Printf("R: %v\n", r)
}
