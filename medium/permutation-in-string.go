package medium

func CheckInclusion(s1 string, s2 string) bool {
	l, count := 0, [26]int{}

	for i := range s1 {
		count[s1[i]-'a']++
	}

	for r := 0; r < len(s2); r++ {
		count[s2[r]-'a']--

		if count == [26]int{} {
			return true
		}

		if r-l+1 == len(s1) {
			count[s2[l]-97]++
			l++
		}
	}
	return false
}
