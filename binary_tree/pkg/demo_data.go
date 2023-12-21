package pkg

func BackExampleData(mode bool) *TreeNode {
	if mode {
		return TrueExample()
	} else {
		return FalseExample()
	}
}

func TrueExample() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
}

func FalseExample() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
}
