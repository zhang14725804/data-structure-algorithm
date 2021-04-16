/*
	迭代方式
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 当前位置
	cur := head
	// 前一个位置
	prev := head
	// careful 【cur != nil】 而不是【cur.Next != nil】
	for cur != nil {
		// 占位符。存储next指针
		next := cur.Next
		// 改变当前节点指针指向（反转嘛） 😅😅😅
		cur.Next = prev
		// 先移动prev，再移动next
		prev = cur
		// 移动cur指针
		cur = next
	}
	// careful 😅😅😅 【head.Next】 切断循环链
	head.Next = nil
	// 最后返回prev
	return prev
}

/*
	方法2：递归方式 genius
*/
func reverseList(head *ListNode) *ListNode {
	// careful： base case 😅😅😅
	if head == nil || head.Next == nil {
		return head
	}
	// 反转head之后的节点
	tail := reverseList(head.Next)
	// 链接head和已反转的部分 😅😅😅
	head.Next.Next = head
	// 切断原来的链接
	head.Next = nil
	// 返回反转之后的链表 😅😅😅
	return tail
}