package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var cycle *ListNode
	cycle = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  4,
				Next: cycle,
			},
		},
	}
	input := &ListNode{
		Val:  3,
		Next: cycle,
	}

	fmt.Println(hasCycle(input))
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil {
		if fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}

	return false
}
