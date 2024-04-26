package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	sentinel := &ListNode{Val: 0}
	sentinel.Next = head

	pred := sentinel
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			pred.Next = head.Next
		} else {
			pred = pred.Next
		}
		head = head.Next
	}
	return sentinel.Next
}
