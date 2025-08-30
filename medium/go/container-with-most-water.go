package main

import "fmt"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	area := 0

	for l < r {
		s := min(height[l], height[r]) * (r - l)

		area = max(area, s)

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return area
}

func main() {
	r := maxArea([]int{1, 1})
	fmt.Printf("R: %v\n", r)
}
