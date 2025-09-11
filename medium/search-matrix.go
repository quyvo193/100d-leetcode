package medium

import "fmt"

func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix)-1, len(matrix[0])-1
	if target < matrix[0][0] || target > matrix[m][n] {
		return false
	}

	l, r := 0, m

	for l < r {
		m := l + (r-l)/2
		fmt.Println("m", m, l, r)

		if matrix[m][n] == target {
			return true
		} else if matrix[m][0] > target {
			if matrix[m][n] <= target {
				return search(matrix[m], target)
			}

			r = m - 1
		} else {
			if matrix[m][n] >= target {
				return search(matrix[m], target)
			}
			l = m + 1
		}
	}

	return search(matrix[l], target)
}

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	fmt.Println("nums", nums)
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
