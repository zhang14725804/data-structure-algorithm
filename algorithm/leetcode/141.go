/*
	给定一个链表，判断链表中是否有环。

	方法1：快慢指针
	方法2：遍历链表，并且把遍历过的节点用 HashSet 存起来，如果遍历过程中又遇到了之前的节点就说明有环。如果到达了 null 就说明没有环
*/
func hasCycle1(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	hash := make(map[int]bool)

	for head != nil {
		hash[head.Val] = true
		head = head.Next
		if _, ok := hash[head.Val]; ok {
			return true
		}
	}
	return false
}