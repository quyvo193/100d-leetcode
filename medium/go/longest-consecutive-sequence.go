package main

import "fmt"

func LongestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	max := 0
	for i := range len(nums) {
		m[nums[i]] = struct{}{}
	}

	for i := range len(nums) {
		consecutive := 0
		x := nums[i]
		for {
			_, next := m[x]
			if !next {
				break
			}

			consecutive++
			delete(m, x)
			x++
		}

		x = nums[i]
		for {
			_, next := m[x-1]
			if !next {
				break
			}

			consecutive++

			delete(m, x-1)
			x--
		}

		if consecutive > max {
			max = consecutive
		}
	}

	return max
}

func main() {
	r := LongestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	fmt.Printf("R: %v\n", r)
}
