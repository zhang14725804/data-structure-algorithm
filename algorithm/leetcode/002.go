package leetcode

/*
	给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

	您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	示例：

	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 0 -> 8
	原因：342 + 465 = 807

	链表，用golang表达链表
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// Leetcode002 第二题
func Leetcode002(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	// 进位
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		curSum := carry
		if l1 != nil {
			curSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curSum += l2.Val
			l2 = l2.Next
		}
		carry = curSum / 10
		//
		temp.Next = &ListNode{Val: curSum % 10}
		temp = temp.Next
	}
	return result.Next
}
