package main

import (
	"fmt"
)

func main() {

	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//Output: [[7,4,1],[8,5,2],[9,6,3]]
	rotate(input)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]         // [0][0] = [2][0]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j] //[2][0] = [2][2]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i] //[2][2] = [0][2]
			matrix[j][n-1-i] = temp                 //[0][2] = [0][0]

		}

	}

	fmt.Println(matrix)
}
