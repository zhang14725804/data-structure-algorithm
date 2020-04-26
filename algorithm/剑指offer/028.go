/*
	在O(1)时间删除链表结点（没有告诉head头节点😅）
*/
type ListNode struct {
	Val int
	Next *ListNode
}
// 巧妙
func deleteNode(node *ListNode)  {
	// 用下一个节点的值覆盖当前节点的值
	node.Val = node.Next.Val
	// 删除下一个节点
	node.Next = node.Next.Next
}