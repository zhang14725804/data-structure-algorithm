/*
	定义一个函数，输入一个链表的头结点，反转该链表并输出反转后链表的头结点。
	todo:测试不通过：输入[]，输出[0]
*/

// 考点：使用一个临时变量记录前序节点
func reverseList(head *ListNode) *ListNode {
	pre := &ListNode{}
	// cur := head
	for head != nil {
		// 缓存下个节点
		next := head.Next
		// todo：下面两行我硬是没看懂😅
		// 遍历下一个节点？
		head.Next = pre
		// 当前节点是下一个节点的pre节点
		pre = head
		// 遍历下一个节点
		head = next
	}
	return pre
}