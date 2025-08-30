package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	return [][]int{}
}

func main() {
	r := twoSum([]int{2, 3, 4}, 6)
	fmt.Printf("R: %v\n", r)
}
