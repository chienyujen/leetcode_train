package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法1
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}

	return pre
}

// 解法2
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre *ListNode
	curr, peer := head, head.Next

	for peer != nil {
		curr.Next = pre
		pre = curr
		curr = peer
		peer = peer.Next
	}

	curr.Next = pre
	return curr
}

// 解法3
func reverseList3(head *ListNode) *ListNode {
	return reverseListRecursively(head)
}

func reverseListRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func main() {
	// Example usage:
	// Initialize a linked list: 1->2->3->4->5
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := reverseList3(head)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
	// Output: 5->4->3->2->1
}
