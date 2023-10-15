package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 這種題目我是真的不會, 參考了chatgpt 才知道怎麼寫...

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 建立一個 dummy 節點，使其指向 linked list 的開頭
	dummy := &ListNode{}
	dummy.Next = head
	// 建立兩個指針，都從 dummy 開始
	first, second := dummy, dummy

	// 將 first 指針往前移動 n+1 個節點
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// 移動 first 指針到結尾，同時移動 second 指針
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// 刪除 second 指針後面的節點
	second.Next = second.Next.Next

	return dummy.Next
}

func main() {
	// Example usage:
	// Initialize a linked list: 1->2->3->4->5
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result := removeNthFromEnd(head, 2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
	// Output: 1->2->3->5
}
