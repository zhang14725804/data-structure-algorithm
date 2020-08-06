/*
	合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

	思路：两两合并
	思路：优先队列（todo：又是优先队列）
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 0 {
		return nil
	}
	// 两两合并
	head := mergeTwoLists(lists[0], lists[1])
	for i := 2; i < len(lists); i++ {
		head = mergeTwoLists(head, lists[i])
	}
	return head
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	// 缓存头节点
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
		// 遍历下个节点
		dummy = dummy.Next
	}
	// 处理剩余部分
	if l1 != nil {
		dummy.Next = l1
	}
	if l2 != nil {
		dummy.Next = l2
	}
	return head.Next
}