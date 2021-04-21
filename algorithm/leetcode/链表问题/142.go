/*
	142. Linked List Cycle II

	快慢指针，开始的时候每次fast走两步slow走一步，相遇之后，【slow】的指针从开头继续走，快慢指针步调一致一起移动，相遇点即为入环点
*/
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	// 快慢指针找到第一次相遇的节点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 相遇
		if fast == slow {
			break
		}
	}
	// 没有环的情况 😄😄
	if fast == nil || fast.Next == nil {
		return nil
	}
	// 【slow】或者【fast】😄😄😄😄 回到起点
	fast = head
	slow = head
	// 找到再次相遇的节点
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}