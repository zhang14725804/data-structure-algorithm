/*
	给定一个链表，判断链表中是否有环。
*/

/*
	方法1：快慢指针 平平无奇
*/
func hasCycle1(head *ListNode) bool {
	fast, slow := head, head
	// 😅😅😅  同时判断 fast和fast.Next 不为nil
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
	方法2：hash
	遍历链表，并且把遍历过的节点用 HashSet 存起来，如果遍历过程中又遇到了之前的节点就说明有环。
	如果到达了 null 就说明没有环
*/
func hasCycle(head *ListNode) bool {
    hash:=make(map[*ListNode]struct{})
    for head!=nil{
        if _,ok:=hash[head];ok{
            return true
		}
		// 先存起来，再遍历下个节点，不能搞反了
        hash[head]=struct{}{}
        head=head.Next
    }
    return false
}