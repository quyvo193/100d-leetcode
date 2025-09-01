package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, t := range tokens {
		switch t {
		case "+":
			f, s := pop(&stack), pop(&stack)
			stack = append(stack, s+f)
			break
		case "-":
			f, s := pop(&stack), pop(&stack)
			stack = append(stack, s-f)
			break
		case "*":
			f, s := pop(&stack), pop(&stack)
			stack = append(stack, s*f)
			break
		case "/":
			f, s := pop(&stack), pop(&stack)
			stack = append(stack, s/f)
			break
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

func main() {
	fmt.Println("result: ", evalRPN([]string{"3", "-4", "+"}))
}
