/*
	链表中倒数第k个节点
	倒数第k个节点就是第n-k+1个节点
*/

type ListNode struct {
	Val int
	Next *ListNode
}
func findKthToTail(head *ListNode, k int) *ListNode {
	// 链表长度
	n := 0
	p:=head
	for p!=nil{
		n++
		p = p.Next
	}
	// 边界问题
	if k > n{
		return nil
	}
	// 找倒数第n个节点
	p = head
	for i:=0;i<n-k;i++{
		p = p.Next
	}
	return p
}