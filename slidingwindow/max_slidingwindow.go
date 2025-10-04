// hard
package slidingwindow

func MaxSlidingWindow(nums []int, k int) []int {
	res, queue := []int{}, []int{}

	for i := 0; i < k; i++ {
		if len(queue) == 0 {
			queue = append(queue, nums[i])
			continue
		}

		for nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}

	}

	l := 0
	r := k

	for r < len(nums) {
		for nums[r] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}

		l++
		r++
	}

	return res
}
