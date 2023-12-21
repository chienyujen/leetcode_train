package main

import (
	"leetcode/binary_tree/pkg"
)

func main() {
	bt := &pkg.TreeNode{}
	isValidBST(bt)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *pkg.TreeNode) bool {
	if root == nil {
		return true
	}
	vals := &[]int{}
	inorder(root, vals)
	for i := 0; i < len(*vals)-1; i++ {
		if (*vals)[i] >= (*vals)[i+1] {
			return false
		}
	}

	return true
}

func inorder(root *pkg.TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, vals)
	*vals = append(*vals, root.Val)
	inorder(root.Right, vals)
}
