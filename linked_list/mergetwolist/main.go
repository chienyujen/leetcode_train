package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	list1 := &ListNode{1, &ListNode{1, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	list2 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}

	result := mergeTwoLists2(list1, list2)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	}

	temp1, temp3 := list1, list1
	temp2, temp4 := list2, list2
	i, j, n := 0, 0, 0
	for temp1 != nil {
		temp1 = temp1.Next
		i++
	}
	for temp2 != nil {
		temp2 = temp2.Next
		j++
	}

	n = i + j

	dummy := &ListNode{}
	curr := dummy
	for s := 0; s <= n; s++ {
		if temp3 == nil {
			curr.Next = temp4
			break
		} else if temp4 == nil {
			curr.Next = temp3
			break
		}

		if temp3.Val <= temp4.Val {
			curr.Next = temp3
			temp3 = temp3.Next
		} else {
			curr.Next = temp4
			temp4 = temp4.Next
		}
		curr = curr.Next
	}
	return dummy.Next
}

func mergeTwoLists2(list1, list2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	return dummy.Next
}
