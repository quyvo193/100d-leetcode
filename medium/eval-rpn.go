package medium

import (
	"strconv"
)

func EvalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, t := range tokens {
		switch t {
		case "+":
			f, s := pop(&stack), pop(&stack)
			stack = append(stack, s+f)
		case "-":
			f, s := pop(&stack), pop(&stack)
			stack = append(stack, s-f)
		case "*":
			f, s := pop(&stack), pop(&stack)
			stack = append(stack, s*f)
		case "/":
			f, s := pop(&stack), pop(&stack)
			stack = append(stack, s/f)
		default:
			num, _ := strconv.Atoi(t)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func pop(s *[]int) (val int) {
	n := len(*s)

	if n == 0 {
		return 0
	}

	val = (*s)[n-1]
	*s = (*s)[:n-1]

	return val
}
