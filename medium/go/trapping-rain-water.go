package main

import "fmt"

func trap(height []int) int {
	l, r := 0, len(height)-1
	ml, mr := height[l], height[r]
	trappedWater := 0
	for r < l {
		ml = max(height[l], ml)
		mr = max(height[r], mr)

		if ml-height[l] > 0 {
			trappedWater += ml - height[l]
		}

		if mr-height[r] > 0 {
			trappedWater += mr - height[r]
		}

		if ml <= mr {
			l++
		} else {
			r--
		}
	}

	return trappedWater
}

func main() {
	r := trap([]int{4, 2, 0, 3, 2, 5})
	fmt.Printf("R: %v\n", r)
}
