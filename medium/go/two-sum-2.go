package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}

	return []int{i + 1, j + 1}
}

func main() {
	r := twoSum([]int{2, 3, 4}, 6)
	fmt.Printf("R: %v\n", r)
}
