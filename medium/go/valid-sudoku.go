package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)

	for i := range 9 {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for x, r := range board {
		for y, c := range r {
			if c == '.' {
				continue
			}

			squareIndex := x/3*3 + y%3

			if rows[x][c] || cols[y][c] || squares[squareIndex][c] {
				return false
			}

			rows[x][c] = true
			cols[y][c] = true
			squares[squareIndex][c] = true
		}
	}

	return true
}

func main() {
	fmt.Println((8 / 3) + (3 / 3))
}
