// Leetcode083 Remove Duplicates from Sorted List
func Leetcode083(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	// 这个指针移到下一个点，是个什么操作，不太理解
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
