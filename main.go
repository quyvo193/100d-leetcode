package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	p := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	// fmt.Println("r: ", easy.Search([]int{1}, 2))
	fmt.Println("r: ", medium.SearchMatrix(p, 11))
}
