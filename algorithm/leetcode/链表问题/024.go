/*
	24.Swap Nodes in Pairs
	1->2->3->4
	2->1->4->3
*/
/*
	方法1：迭代 😅😅😅
	question
*/
func swapPairs(head *ListNode) *ListNode {
	// 虚拟头结点
	dummy := &ListNode{}
	dummy.Next = head

	p := dummy
	for p.Next != nil && p.Next.Next != nil {
		a := p.Next
		b := p.Next.Next
		// 😅😅😅
		// 改变头节点
		p.Next = b
		// 改变头节点指向
		a.Next = b.Next
		// 改变下个节点指向
		b.Next = a
		// 遍历下个节点
		p = a
	}
	return dummy.Next
}

/*
	方法2：递归 😅😅😅
	question
	思路：将链表翻转转化为一个子问题，然后通过递归方式依次解决
*/
func swapPairs(head *ListNode) *ListNode {
	return helper(head)
}

func helper(head *ListNode) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	// 保存下一阶段的头指针
	nextHead := head.Next.Next
	// 反转当前阶段指针
	next := head.Next
	next.Next = head
	head.Next = helper(nextHead)
	// 返回next
	return next
}