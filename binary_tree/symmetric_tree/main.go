package main

import (
	"fmt"
	"leetcode/binary_tree/pkg"
)

func main() {
	runMode := true
	res := isSymmetric(pkg.BackExampleData(runMode))
	fmt.Println(res)
}

func isSymmetric(root *pkg.TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(t1, t2 *pkg.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)

}
