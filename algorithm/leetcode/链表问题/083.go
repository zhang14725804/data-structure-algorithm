/*
	给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
	思路：快慢指针
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	// 移动快慢指针
	for fast != nil {
		// 不重复的情况
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 😅 careful 断开后续重复数据。排除末尾有重复的情况（[1,1,2,3,3]）
	slow.Next = nil
	return head
}

/*
	常规操作
*/
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := head
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			// 过滤重复值
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return dummy
}