/*

 */
func partition(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	sHead, lHead := small, large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	// 😅 断尾巴
	large.Next = nil
	// 😅 small链接large
	small.Next = lHead.Next
	return sHead.Next
}