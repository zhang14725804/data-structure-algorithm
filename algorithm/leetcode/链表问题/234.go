/*
	请判断一个链表是否为回文链表。
	你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？ 😅😅😅
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
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转尾部【slow】指向的部分 😄😄😄
	tail := reverseList(slow)
	// 头尾两部逐一分比较val
	for tail != nil {
		if tail.Val != head.Val {
			return false
		}
		tail = tail.Next
		head = head.Next
	}
	return true
}

/*
	方法2： 双指针法

	对于单链表，无法直接倒序遍历，可以造一条新的反转链表，可以利用链表的后序遍历，也可以用栈结构倒序处理单链表。
	借助二叉树后序遍历的思路，不需要显式反转原始链表也可以倒序遍历链表
	实际上就是把链表节点放入一个栈，然后再拿出来，这时候元素顺序就是反的
	(question 没理解这思想 😅😅😅😅😅😅 那个大神讲的忘记了)
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