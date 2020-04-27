/*
	 删除链表中重复的节点
*/
type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplication(head *ListNode) *ListNode {
	// 虚拟头节点（有可能删除头节点）
	dummy := &ListNode{}
	dummy.Next = head
	// 双指针方法（todo：没看懂😅）
	p := dummy
	for p.Next != nil{
		q := p.Next
		// 重复元素
		for q!=nil && p.Next.Val == q.Val {
			q = q.Next
		}
		if p.Next.Next == q {
			p = p.Next
		}else{
			p.Next = q
		}
	}
	return dummy.Next
}