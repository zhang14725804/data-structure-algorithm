/*
	给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。
	请你将 list1 中第 a 个节点到第 b 个节点删除，并将list2 接在被删除节点的位置。
*/
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head, dummy := list1, list1
	begin, end := &ListNode{}, &ListNode{}
	for list1.Next != nil {
		a--
		b--
		list1 = list1.Next
		// 第a个节点的前一个节点
		if a == 1 {
			begin = list1
		}
		// 第b个节点
		if b == 0 {
			end = list1.Next
			break
		}
	}
	for head.Next != nil {
		// 拼接list2
		if head.Val == begin.Val {
			head.Next = list2
		}
		head = head.Next
	}
	// 拼接剩下的部分
	head.Next = end
	return dummy
}