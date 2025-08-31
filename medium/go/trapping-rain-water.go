package main

import "fmt"

func trap(height []int) int {

	l, r := 0, 1
	water := 0
	for {
		if height[l] > height[r] {
			r++
		}

		if height[l] <= height[r] {
			water += calculateWater(height[l:r])
			l = r
			r = l + 1
		}

		if r == len(height)-1 {
			break
		}
	}

	return water
}

func calculateWater(sub []int) int {
	subWater := 0
	height := min(sub[0], sub[len(sub)-1])

	for i := 1; i < len(sub)-2; i++ {
		subWater += height - sub[i]
	}

	return subWater
}

func main() {
	r := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Printf("R: %v\n", r)
}
