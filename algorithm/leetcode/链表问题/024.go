/*
	24.Swap Nodes in Pairs

	p.Next = b
	a.Next = b.Next
	b.Next = a
	p = a
*/
func Leetcode024(head *ListNode) *ListNode {
	// 虚拟头结点
	newHead := &ListNode{}
	newHead.Next = head

	p := newHead
	for p.Next != nil && p.Next.Next != nil {
		a := p.Next
		b := p.Next.Next

		p.Next = b
		a.Next = b.Next
		b.Next = a
		p = a
	}
	return newHead.Next
}
