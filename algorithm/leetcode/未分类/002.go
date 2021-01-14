/*
	给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
	如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
	您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	思路
*/
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	// 虚拟头节点
	dummy := &ListNode{}
	cur := dummy
	// 进位
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		// 当前和
		curSum := carry
		// 遍历两个链表
		if l1 != nil {
			curSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curSum += l2.Val
			l2 = l2.Next
		}
		// 进位
		carry = curSum / 10
		// curSum % 10当前位的数值
		cur.Next = &ListNode{Val: curSum % 10}
		// cur前进一步
		cur = cur.Next
	}
	return dummy.Next
}
