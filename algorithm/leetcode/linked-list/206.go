package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	206. Reverse Linked List（链表反转）
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
		// 改变指针指向
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
