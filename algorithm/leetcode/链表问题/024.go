/*
	方法1：迭代 😅😅😅
	question 😅😅😅
*/
func swapPairs(head *ListNode) *ListNode {
	// 😅 虚拟头结点
	dummy := &ListNode{}
	dummy.Next = head

	p := dummy
	for p.Next != nil && p.Next.Next != nil {
		first := p.Next
		second := p.Next.Next
		// 😅 反转头
		p.Next = second
		// 😅 反转第一个
		first.Next = second.Next
		// 😅 反转第二个
		second.Next = first
		// 😅 遍历下个节点，此时first指向下一组
		p = first
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
	// 😅 保存下一阶段的头指针
	nextHead := head.Next.Next
	// 反转第二个
	next := head.Next
	next.Next = head
	// 😅 反转下个阶段
	head.Next = helper(nextHead)
	// 返回next
	return next
}