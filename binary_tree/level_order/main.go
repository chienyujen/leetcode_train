package main

import (
	"fmt"
	"leetcode/binary_tree/pkg"
)

func main() {
	runMode := true

	res := levelOrder(pkg.BackExampleData(runMode))

	fmt.Println(res)
}

func levelOrder(root *pkg.TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := []*pkg.TreeNode{root}

	for len(queue) > 0 {
		level := []int{}
		currentSize := len(queue)
		for i := 0; i < currentSize; i++ {
			current := queue[0]
			queue = queue[1:]
			level = append(level, current.Val)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
