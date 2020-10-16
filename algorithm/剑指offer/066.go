/*
	两个链表，找出它们的第一个公共结点
	todo：思路精巧😄
*/
func findFirstCommonNode(headA *ListNode, headB *ListNode) *ListNode {
	p, q := headA, headB
	for q != p {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}
		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}
	return q
}