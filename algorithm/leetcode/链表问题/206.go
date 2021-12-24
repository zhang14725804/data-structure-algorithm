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
	// careful 😅😅😅 【cur != nil】 而不是【cur.Next != nil】
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
	// 【head.Next】 切断循环链 careful 😅😅😅
	head.Next = nil
	// 最后返回prev
	return prev
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 前一个指针
	var prev *ListNode = nil
	// 当前指针
	cur := head
	// 缓存下一个指针
	var next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

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

/*
	方法3 2021-12-24
*/
func reverseList(head *ListNode) *ListNode {
	// 用一个slice存储每个节点
	nodeList := make([]*ListNode, 0)
	for head != nil {
		nodeList = append(nodeList, head)
		head = head.Next
	}
	dummy := &ListNode{}
	nh := dummy
	// 倒叙遍历
	for i := len(nodeList) - 1; i >= 0; i-- {
		// 加入新链表
		dummy.Next = nodeList[i]
		// 切断原指针
		nodeList[i].Next = nil
		dummy = dummy.Next
	}
	return nh.Next
}