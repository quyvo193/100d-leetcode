// medium problem
package twopointers

func MaxArea(height []int) int {
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
