/*
	将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/

/*
	方法1：迭代
*/ 
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// var dummy *ListNode  // nil 未初始化
	dummy := &ListNode{} // &{0 <nil>} 初始化的情况
	// 缓存头节点 😅😅😅
	head := dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			dummy.Next = l2
			// 遍历下个节点
			l2 = l2.Next
		} else {
			dummy.Next = l1
			l1 = l1.Next
		}
		// 遍历下个节点 😅😅😅
		dummy = dummy.Next
	}
	// 处理剩余部分
	if l1 != nil {
		dummy.Next = l1
	}
	if l2 != nil {
		dummy.Next = l2
	}
	// 返回 😅😅😅
	return head.Next
}

/*
	方法2：递归 😅😅😅
*/ 
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// base case
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// genius
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} 
	l2.Next = mergeTwoLists(l2.Next, l1)
	return l2
}