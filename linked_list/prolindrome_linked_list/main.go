package main

import (
	"fmt"
	"runtime"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	fmt.Println(isPalindrome2(input))
}

func isPalindrome(head *ListNode) bool {
	rec := "" // 反向字串
	seq := "" // 正向字串

	for head != nil {
		seq = fmt.Sprintf("%s%d", seq, head.Val)
		rec = fmt.Sprintf("%d%s", head.Val, rec)
		head = head.Next
	}
	if rec != seq {
		return false
	}

	return true
}

func isPalindrome2(head *ListNode) bool {
	defer runtime.GC()
	middle := findMiddle(head)
	rNode := reverseNode(middle)

	for rNode != nil {
		if head.Val != rNode.Val {
			return false
		}
		head = head.Next
		rNode = rNode.Next
	}

	return true
}

func findMiddle(node *ListNode) *ListNode {
	fast := node
	slow := node

	for (fast != nil) && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseNode(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}

	var prev *ListNode
	cur := node

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
