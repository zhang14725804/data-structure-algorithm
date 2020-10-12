/*
	删除链表中等于给定值 val 的所有节点。
*/

// 方法1：遍历
func removeElements(head *ListNode, val int) *ListNode {
	// 虚拟头节点，方便考虑头节点
	dummy := &ListNode{}
	dummy.Next = head
	p := dummy
	for p.Next != nil {
		//
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}

// 方法2：递归
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return removeElements(head.Next, val)
	}
	head.Next = removeElements(head.Next, val)
	return head
}