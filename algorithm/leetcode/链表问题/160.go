/*
	思路1：
	1. 两个指针first,second
	2. 一个从a开始遍历，走完a之后再从b开始走；
	3. 一个从b开始遍历，走完b之后再从a开始走
	4. 而这相遇的点就是相交的起始点。
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	first := headA
	second := headB

	for first != second {
		// 遍历a
		if first != nil {
			first = first.Next
		} else {
			// 指针指向b，再遍历b
			first = headB
		}
		// 遍历b
		if second != nil {
			second = second.Next
		} else {
			// 指针指向a，再遍历a
			second = headA
		}
	}
	return first
}
