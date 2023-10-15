package main

import (
	"fmt"
)

func main() {
	// fmt.Println(1 / 3)
	data := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println(isValidSudoku(data))
}

func isValidSudoku(board [][]byte) bool {
	var row, col, subBox [9][9]bool
	for seq1, rowData := range board {
		for seq2, colData := range rowData {
			if colData == '.' {
				continue
			}
			num := int(colData) - 49
			k := seq1/3*3 + seq2/3
			if row[seq1][num] || col[seq2][num] || subBox[k][num] {
				return false
			}

			row[seq1][num], col[seq2][num], subBox[k][num] = true, true, true
		}
	}
	return true
}
