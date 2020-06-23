/*
	给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
	k 是一个正整数，它的值小于或等于链表的长度。
	如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

	注意：

	你的算法只能使用常数的额外空间。
	你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

	todo：不太懂
*/

type ListNode struct {
    Val int
    Next *ListNode
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 虚拟头节点
	dummy:=&ListNode{}
	dummy.Next = head
	cur := dummy

	for cur!=nil{
		// 判断是否够k个节点
		s:=0
		for i:=cur.Next;i!=nil;i = i.Next{
			s++
		}
		if s<k{
			break
		}

		s = 0
		// 双指针记录相邻节点
		a := cur.Next
		b := a.Next
		// 反转k-1次
		for s < k-1{
			s++
			// 真乱😅
			c := b.Next
			b.Next = a
			a = b
			b = c
		}

		// 更乱了尴尬
		p := cur.Next
		cur.Next.Next = b
		cur.Next = a
		cur = p
	}
	return dummy.Next
}