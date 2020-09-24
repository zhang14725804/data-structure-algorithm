/*
	请判断一个链表是否为回文链表。
*/

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 快慢指针找到中间位置
	slow, fast := head, head
	// 这样不用考虑奇数偶数
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转尾部
	tail := reverseList(slow)
	// 头尾两部分比较
	for tail != nil {
		if tail.Val != head.Val {
			return false
		}
		tail = tail.Next
		head = head.Next
	}
	return true
}

// 链表反转
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 当前节点、前一个节点
	cur, prev := head, head
	for cur != nil {
		// 缓存next指针
		next := cur.Next
		// 改变当前节点指针指向
		cur.Next = prev
		// 移动prev指针
		prev = cur
		// 移动当前指针
		cur = next
	}
	// 切断循环链
	head.Next = nil
	// 返回 prev
	return prev
}