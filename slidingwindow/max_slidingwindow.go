// hard
package slidingwindow

import (
	"fmt"
	"math"
)

func findFirstAndSecondLargest(arr []int, maxMap *map[int]int) (int, int) {
	if len(arr) < 2 {
		return arr[0], math.MinInt
	}

	firstLargest := math.MinInt
	secondLargest := math.MinInt

	for _, num := range arr {
		if num > firstLargest {
			secondLargest = firstLargest
			firstLargest = num
			(*maxMap)[firstLargest]++
		} else if num > secondLargest && num != firstLargest {
			secondLargest = num
		}
	}

	return firstLargest, secondLargest
}

func MaxSlidingWindow(nums []int, k int) []int {
	l, r := 0, k-1

	maxMap := make(map[int]int)
	first, second := findFirstAndSecondLargest(nums[l:r+1], &maxMap)
	res := []int{first}
	r++
	l++

	for r < len(nums) {
		fmt.Println(first, second, maxMap)
		if first == nums[l-1] {
			maxMap[first]--
		}
		if nums[r] >= first {
			second = first
			first = nums[r]
			maxMap[first]++
			res = append(res, nums[r])
		} else if maxMap[first] > 0 {
			res = append(res, first)
			second = max(second, nums[r])
		} else {
			if second < nums[r] {
				first = nums[r]
				maxMap[first]++
				res = append(res, first)
			} else {
				first, second = findFirstAndSecondLargest(nums[l:r+1], &maxMap)
				res = append(res, first)
			}
		}

		r++
		l++
	}

	return res
}
