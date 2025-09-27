package arrayhashing

func TopKFrequent(nums []int, k int) []int {
	duplicatedMap := make(map[int]int)
	freq := make([][]int, len(nums))
	res := []int{}

	for _, v := range nums {
		duplicatedMap[v]++
	}

	for key, val := range duplicatedMap {
		freq[val-1] = append(freq[val-1], key)
	}

	for i := len(freq) - 1; i >= 0; i-- {
		if len(freq[i]) > 0 {
			for _, v := range freq[i] {
				res = append(res, v)
				k--
			}
			if k == 0 {
				break
			}
		}
	}

	return res
}
