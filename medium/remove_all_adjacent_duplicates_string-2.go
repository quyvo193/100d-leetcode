package medium

func RemoveDuplicates(s string, k int) string {
	stack := []byte{}

	for i := range s {
		count := 1
		for len(s) > 0 && stack[len(s)-count] == s[i] {
			count++

			if count == k {
				stack = stack[:len(s)-k]
			}
		}

		stack = append(stack, s[i])
	}

	return string(s)
}
