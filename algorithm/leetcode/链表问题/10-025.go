/*
	给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
	k 是一个正整数，它的值小于或等于链表的长度。
	如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

	注意：
	你的算法只能使用常数的额外空间。
	你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

/*
	方法1：整体迭代实现 （难懂）
*/
func reverseKGroup1(head *ListNode, k int) *ListNode {
	// 😅😅😅 虚拟头节点
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy

	for cur != nil {
		// 😅 判断是否够k个节点
		cLen := 0
		for i := cur.Next; i != nil; i = i.Next {
			cLen++
		}
		// 不足k个，无需反转；
		if cLen < k {
			break
		}

		cLen = 0
		// 😅😅😅 双指针记录相邻节点
		prev := cur.Next
		prevNext := prev.Next
		// 😅 反转k-1次
		for cLen < k-1 {
			cLen++
			// 缓存下个节点
			temp := prevNext.Next
			// 😅 反转指针
			prevNext.Next = prev
			// prev向前走一步
			prev = prevNext
			// prevNext向前走一步
			prevNext = temp
		}

		// 缓存下个节点
		next := cur.Next
		// 😅 question 反转指针
		cur.Next.Next = prevNext
		// 😅 question
		cur.Next = prev
		// next 向前走一步
		cur = next
	}
	return dummy.Next
}

/*
	方法2：整体递归+部分迭代
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	// 😅😅😅 区间[left,right)包含k个待反转元素
	left, right := head, head
	for i := 0; i < k; i++ {
		// base case： 不足k个，无需反转；
		if right == nil {
			// 😅😅😅 返回head
			return head
		}
		right = right.Next
	}
	// 😅😅😅 反转[left,right)之间的元素
	dummy := reverse(left, right)
	// 😅😅😅 递归反转后续链表，并链接
	left.Next = reverseKGroup(right, k)
	return dummy
}

// 反转[left,right)之间的元素 😅
func reverse(left, right *ListNode) *ListNode {
	cur := left
	var prev, next *ListNode
	for cur != right {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}