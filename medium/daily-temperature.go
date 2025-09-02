package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		if i == 0 || temperatures[i] < temperatures[i-1] {
			for j := i + 1; j < len(temperatures); j++ {
				if temperatures[i] < temperatures[j] {
					res[i] = j - i
					break
				}
			}
		} else {
			if i > 0 && res[i-1] == 0 {
				res[i] = 0
				continue
			}

			for j := i + res[i-1] - 1; j < len(temperatures); j++ {
				if temperatures[i] < temperatures[j] {
					res[i] = j - i
					break
				}
			}
		}

	}

	return res
}

func main() {
	fmt.Println("result: ", dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
