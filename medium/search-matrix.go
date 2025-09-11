package medium

func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix)-1, len(matrix[0])-1
	if target < matrix[0][0] || target > matrix[m][n] {
		return false
	}

	l, r := 0, m
	for l < r {
		mid := (l + r) / 2

		if matrix[mid][n] < target {
			l = mid + 1
		} else if matrix[mid][0] > target {
			r = mid - 1
		} else {
			break
		}
	}

	if l > r {
		return false
	}

	return search(matrix[(l+r)/2], target)
}

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return true
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}

	}

	return false
}
