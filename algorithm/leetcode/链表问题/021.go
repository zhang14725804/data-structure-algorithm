/*
	方法1：迭代（这种方法我会）
	1. 当l1不为空且l2不为空时，依次遍历两个链，比较大小
	2. 处理剩余部分
	3. 比较难理解的是dummy部分
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
	方法2：递归 （喜欢这种方法）😅😅😅
	1. 定义递归出口，l1或者l2位空时
	2. 都不为空，那个小以那个为头

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