package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	input := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  9,
					Next: &ListNode{},
				},
			},
		},
	}

	deleteNode(input.Next.Next)

	fmt.Printf("%+v", input.Next.Next)
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
