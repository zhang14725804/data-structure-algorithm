package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	206. Reverse Linked List
*/
func Leetcode206(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	a := head
	b := head.Next
	for b != nil {
		c := b.Next
		b.Next = a
		a = b
		b = c
	}
	head.Next = nil
	return a
}
