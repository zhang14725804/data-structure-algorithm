/*
	160. Intersection of Two Linked Lists（找到两个单链表相交的起始节点。）
	思路1：两个指针，一个从a开始遍历，走完a之后再从b开始走；一个从b开始遍历，走完b之后再从a开始走，而这相遇的点就是相交的起始点。
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p := headA
	q := headB

	for p != q {
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
	return p
}
