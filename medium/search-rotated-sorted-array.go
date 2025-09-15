package medium

func Search(nums []int, target int) int {
	minIdx := findMinIndex(nums)

	if target >= nums[0] {
		if minIdx == 0 {
			return binarySearch(nums[0:], target)
		}

		return binarySearch(nums[0:minIdx], target)
	} else {
		pos := binarySearch(nums[minIdx:], target)
		if pos >= 0 {
			return minIdx + pos
		}
	}

	return -1
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if target > nums[m] {
			l = m + 1
		} else {
			r = m + 1
		}
	}

	return -1
}

func findMinIndex(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2

		if nums[m] < nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
