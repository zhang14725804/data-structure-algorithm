package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// Leetcode083 Remove Duplicates from Sorted List
func Leetcode083(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
