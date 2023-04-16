/*
	解题思路：快慢指针
	1. 开始的时候每次fast走两步，slow走一步
	2. 考虑相遇和没有环的情况
	3. 相遇之后，【slow】的指针从开头继续走，快慢指针步调一致一起移动，相遇点即为入环点
*/
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	// 1. 判断fast 😅😅😅
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 相遇
		if fast == slow {
			break
		}
	}
	// 2. 没有环的情况 😄😄
	if fast == nil || fast.Next == nil {
		return nil
	}
	// 4.【slow】或者【fast】😄😄😄😄 回到起点
	fast = head
	// 5.  找到再次相遇的节点
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}