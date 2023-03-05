/*
	思路：快慢指针
	1. slow每次走一步
	2. fast根据是否有重复元素每次都走
	3. 没有重复元素的情况 slow.Next=fast
	4. 否则fast走，slow不走
	5. 最后记得切尾巴 😅😅😅
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	// 移动快慢指针
	for fast != nil {
		// slow只有在不同元素的情况下才走
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		// fast每次都走
		fast = fast.Next
	}
	// 😅 careful 断开后续重复数据。排除末尾有重复的情况（[1,1,2,3,3]）
	slow.Next = nil
	return head
}

/*
	常规操作 😅😅😅
*/
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := head
	for head != nil {
		// 😅 这里不能用for（为什么），注意边界判定
		if head.Next != nil && head.Val == head.Next.Val {
			// 过滤一个重复值
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return dummy
}