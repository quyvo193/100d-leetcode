package hard

func LargestRectangleArea(heights []int) int {
	n := len(heights)
	m := 0

	for i := range n {
		l, r := i, i

		for l > 0 && heights[i] <= heights[l] {
			l--
		}

		if heights[i] > heights[l] {
			l++
		}

		for r < n-1 && heights[i] <= heights[r] {
			r++
		}

		if heights[i] > heights[r] {
			r--
		}

		m = max((r-l+1)*heights[i], m)
	}

	return m
}
