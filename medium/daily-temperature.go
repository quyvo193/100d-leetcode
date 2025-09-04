package medium

type HigherTemp struct {
	idx int
	val int
}

func DailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []HigherTemp{}

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > stack[len(stack)-1].val {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[pop.idx] = i - pop.idx
		}

		stack = append(stack, HigherTemp{
			idx: i,
			val: temp,
		})
	}

	return res
}
