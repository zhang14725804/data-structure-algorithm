package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// Leetcode082 Remove Duplicates from Sorted List II
func Leetcode082(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{}
	temp := newHead

	// todos，不太懂
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			head = head.Next
		} else {
			temp.Next = head
			temp = temp.Next
			head = head.Next
		}
	}
	temp.Next = head
	return newHead.Next
}
