/*
	给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
	k 是一个正整数，它的值小于或等于链表的长度。
	如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

	注意：

	你的算法只能使用常数的额外空间。
	你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

/*
	方法1：迭代实现(question 😅)，看不懂
*/
func reverseKGroup1(head *ListNode, k int) *ListNode {
	// 虚拟头节点
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy

	for cur != nil {
		// 判断是否够k个节点
		s := 0
		for i := cur.Next; i != nil; i = i.Next {
			s++
		}
		if s < k {
			break
		}

		s = 0
		// 双指针记录相邻节点
		a := cur.Next
		b := a.Next
		// 反转k-1次
		for s < k-1 {
			s++
			// 真乱😅
			c := b.Next
			b.Next = a
			a = b
			b = c
		}

		// 更乱了尴尬
		p := cur.Next
		cur.Next.Next = b
		cur.Next = a
		cur = p
	}
	return dummy.Next
}

/*
	方法2：递归+迭代
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间[left,right)包含k个待反转元素
	left, right := head, head
	for i := 0; i < k; i++ {
		// 不足k个，无需反转；base case
		if right == nil {
			return head
		}
		right = right.Next
	}
	// 反转前k个元素
	dummy := reverse(left, right)
	// 递归反转后续链表，并连接起来
	left.Next = reverseKGroup(right, k)
	return dummy
}

// 反转[a,b)之间的元素，左闭右开
func reverse(a, b *ListNode) *ListNode {
	var prev, cur, next *ListNode
	cur, next = a, a
	for cur != b {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}