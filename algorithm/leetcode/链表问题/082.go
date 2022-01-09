/*
	存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字
	链表头结点可能被删除，所以用 dummy node 辅助删除
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	//  😅😅😅 遍历指针
 	cur := dummy

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			// 循环删除重复节点，这里是 for 而不是 if
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			head = head.Next
		} else {
			// 指向新的头节点
			cur.Next = head
			// 移动指针
			cur = cur.Next
			// 移动指针
			head = head.Next
		}
	}
	// 处理最后一个节点
	cur.Next = head
	// 返回跳过头部虚拟节点之后的头节点
	return dummy.Next
}
