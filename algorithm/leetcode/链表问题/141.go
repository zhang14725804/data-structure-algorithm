/*
	方法1：快慢指针 平平无奇
	1. slow每次走一步，fast每次走两步
	2. 若重合，代表有环
	3. 同时判断 fast和fast.Next 不为nil 😅😅😅
*/
func hasCycle1(head *ListNode) bool {
	fast, slow := head, head
	//  😅😅😅
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 比较指针是否相等（不要使用val比较！）
		if fast == slow {
			return true
		}
	}
	return false
}

/*
	方法2：hash 这方法也不错
	1. 遍历链表，并且把遍历过的节点用 HashSet 存起来
	2. 如果遍历过程中又遇到了之前的节点就说明有环。
	3. 如果到达了 null 就说明没有环
*/
func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}
		// 先存起来，再遍历下个节点，不能搞反了
		hash[head] = struct{}{}
		head = head.Next
	}
	return false
}