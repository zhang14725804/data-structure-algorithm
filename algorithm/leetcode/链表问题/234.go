/*
	请判断一个链表是否为回文链表。
	你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

/*
	方法1：
		（1）快慢指针找到中间位置
		（2）反转尾部
		（3）比较头部和反转之后的尾部
*/
func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 快慢指针找到中间位置
	slow, fast := head, head
	// (question 小技巧，但是为啥😅)这样不用考虑奇数偶数
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

/*
	方法2：(question 妙啊！！)
	双指针法
	对于单链表，无法直接倒序遍历，可以造一条新的反转链表，可以利用链表的后序遍历，也可以用栈结构倒序处理单链表。
	借助二叉树后序遍历的思路，不需要显式反转原始链表也可以倒序遍历链表
	实际上就是把链表节点放入一个栈，然后再拿出来，这时候元素顺序就是反的(question 没理解这思想😅)
*/
var left *ListNode

func isPalindrome(head *ListNode) bool {
	left = head
	return traverse(head)
}

func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := traverse(right.Next)
	// 后序遍历
	res = res && right.Val == left.Val
	left = left.Next
	return res
}