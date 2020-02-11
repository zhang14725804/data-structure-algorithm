package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	160. Intersection of Two Linked Lists
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
