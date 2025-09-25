// medium
package arrayhashing

import "math"

func MaxSubArray(nums []int) int {
	m := math.MinInt
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		m = max(m, sum)

		if sum < 0 {
			sum = 0
		}
	}

	return m
}
