package medium

import (
	"math"
)

func MinEatingSpeed(piles []int, h int) int {
	mb := 0

	for _, v := range piles {
		mb = max(mb, v)
	}

	res := mb
	l, r := 1, mb

	for l <= r {
		totalHours := 0
		m := (l + r) / 2

		for i := range piles {
			totalHours += int(math.Ceil(float64(piles[i]) / float64(m)))
		}

		if totalHours <= h {
			r = m - 1
			res = min(res, m)
		} else {
			l = m + 1
		}
	}

	return res
}
