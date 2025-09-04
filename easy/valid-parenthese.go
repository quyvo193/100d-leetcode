package easy

// {([])}
func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
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
