/*
	存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字
	链表头结点可能被删除，所以用 dummy node 辅助删除
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	//  question 这一步什么作用  😅😅😅
	temp := dummy

	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			// 循环删除重复节点
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			head = head.Next
		} else {
			// 指向新的头节点
			temp.Next = head
			// 移动指针
			temp = temp.Next
			// 移动指针
			head = head.Next
		}
	}
	// question 没看懂  😅😅😅
	// dummy全程观看，最后还是正确的
	temp.Next = head
	// careful：返回跳过头部虚拟节点之后的头节点
	return dummy.Next
}
