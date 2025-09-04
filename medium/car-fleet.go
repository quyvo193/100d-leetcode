package medium

import (
	"slices"
)

type car struct {
	p int
	t float32
}

func CarFleet(target int, position []int, speed []int) int {
	cars := make([]car, 0)
	fleet := 1

	for i := range len(position) {
		cars = append(cars, car{p: position[i], t: float32(target-position[i]) / float32(speed[i])})
	}

	slices.SortFunc(cars, func(c1, c2 car) int {
		return c2.p - c1.p
	})

	m := cars[0].t

	for i := range len(cars) {
		car := cars[i]

		if car.t > m {
			fleet++
			m = car.t
		}
	}

	return fleet
}

// 8 6 -> 10
// 2 3
