/*
	142. Linked List Cycle II

	快慢指针，开始的时候一个走两步一个走一步，相遇之后，慢的指针从开头继续走，另一个再相遇点开始走，每次都走一步，一定会在b点相遇
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
	// 没有环的情况
	if fast == nil || fast.Next == nil {
		return nil
	}
	// slow回到起点
	slow = head
	// 找到再次相遇的节点
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}