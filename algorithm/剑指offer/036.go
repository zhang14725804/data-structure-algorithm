/*
	输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍然是按照递增排序的。
	todo
*/
type ListNode struct {
    Val int
    Next *ListNode
}
// 归并算法
func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	// 虚拟头节点
	dummy := &ListNode{}
	cur := dummy
	// 遍历两个链表
	for l1 != nil && l2 !=nil{
		if l1.Val < l2.Val{
			cur.Next = l1
			// 注意：移动指针
			cur = l1
			l1 = l1.Next
		}else{
			cur.Next = l2
			// 注意：移动指针
			cur = l2
			l2 = l2.Next
		}
	}
	// 剩余的部分
	if l1 != nil{
		cur.Next = l1
	}else{
		cur.Next = l2
	}
	// 注意：不是cur.Next
	return dummy.Next
}