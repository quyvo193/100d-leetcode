package main

import (
	"fmt"
	"strings"
)

func isValidSudoku(board [][]byte) bool {
	rows := make(map[byte]bool)
	cols := make(map[byte]bool)
	squares := make(map[string]bool)

	for x, r := range board {
		for y, c := range r {
			if c == '.' {
				continue
			}

			if rows[c] || cols[c] || squares[strings.Join([]string{fmt.Sprint(x % 3), string(y % 3)}, "-")] {
				return false
			}

			rows[c] = true
			cols[c] = true
			squares[strings.Join([]string{string(x % 3), string(y % 3)}, "-")] = true
		}
	}

	return true
}
