/*
	思路1： 0102 依旧没思路
	两个指针，一个从a开始遍历，走完a之后再从b开始走；
	一个从b开始遍历，走完b之后再从a开始走，而这相遇的点就是相交的起始点。
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p := headA
	q := headB

	for p != q {
		// 遍历a
		if p != nil {
			p = p.Next
		} else {
			// 指针指向b，再遍历b
			p = headB
		}
		// 遍历b
		if q != nil {
			q = q.Next
		} else {
			// 指针指向a，再遍历a
			q = headA
		}
	}
	return p
}
