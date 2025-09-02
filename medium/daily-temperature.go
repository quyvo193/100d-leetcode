package main

import "fmt"

type HigherTemp struct {
	idx int
	val int
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, len(temperatures))
	stack := []HigherTemp{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1].val <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			res[i] = stack[len(stack)-1].idx - i
		}

		stack = append(stack, HigherTemp{
			idx: i,
			val: temperatures[i],
		})
	}

	return res
}

func main() {
	fmt.Println("result: ", dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
}
