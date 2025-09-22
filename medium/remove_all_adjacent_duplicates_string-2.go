package medium

type Item struct {
	Char  byte
	Count int
}

func RemoveDuplicates(s string, k int) string {
	stack := []Item{}
	res := []byte{}

	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1].Char == s[i] && stack[len(stack)-1].Count == k-1 {
			stack = stack[:len(stack)-k]
		} else {
			if stack[len(stack)-1].Char == s[i] {
				stack = append(stack, Item{Char: s[i], Count: stack[len(stack)-1].Count + 1})
			} else {
				stack = append(stack, Item{Char: s[i], Count: 1})
			}
		}
	}

	for _, item := range stack {
		res = append(res, item.Char)
	}

	return string(res)
}
