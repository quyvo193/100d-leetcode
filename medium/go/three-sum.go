package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := [][]int{}

	for i := range len(nums) {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
			} else if sum > 0 {
				l++
				for nums[l] == nums[l-1] {
					l++
				}
			} else {
				r--
			}

		}
	}

	return res
}

func main() {
	a := []int{}

	a = append(a, 5)
	// r := twoSum([]int{2, 3, 4}, 6)
	fmt.Printf("R: %v\n", a)
}
