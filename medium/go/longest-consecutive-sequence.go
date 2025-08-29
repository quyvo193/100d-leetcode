package main

import "fmt"

func LongestConsecutive(nums []int) int {
	m := make(map[int]bool)
	max := 0
	for i := range len(nums) {
		m[nums[i]] = true
	}

	for _, num := range nums {
		if m[num-1] {
			continue
		}

		seq, temp := 1, num+1

		for m[temp] {
			seq++
			temp++
		}

		if seq > max {
			max = seq
		}

		if max > len(nums)/2 {
			break
		}
	}

	return max
}

func main() {
	r := LongestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	fmt.Printf("R: %v\n", r)
}
