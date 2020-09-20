/*
	todo:链表反转（迭代或递归）
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 当前位置
	cur := head
	// 前一个位置
	prev := head
	for cur != nil {
		// 占位符。存储next指针
		next := cur.Next
		// 改变当前节点指针指向（反转嘛） 😅
		cur.Next = prev
		// 移动prev指针
		prev = cur
		// 移动cur指针
		cur = next
	}
	// 切断循环链
	head.Next = nil
	return prev
}

// 迭代方式
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	tail := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return tail
}