package medium

import "fmt"

func CheckInclusion(s1 string, s2 string) bool {
	charMap := make(map[byte]int)

	for i := range s1 {
		charMap[s1[i]]++
	}

	fmt.Println(charMap)

	return true
}
