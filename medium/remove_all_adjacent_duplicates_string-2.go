package medium

func RemoveDuplicates(s string, k int) string {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		count := 1

		for len(stack) > 0 && stack[len(stack)-count] == s[i] {
			if count == k-1 {
				stack = stack[:len(stack)-k+1]
				goto endloop
			}

			count++
			if count > len(stack) {
				break
			}
		}

		stack = append(stack, s[i])
	endloop:
	}

	return string(stack)
}
