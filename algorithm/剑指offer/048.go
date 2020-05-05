/*
	在复杂链表中，每个结点除了有一个指针指向下一个结点外，还有一个额外的指针指向链表中的任意结点或者null
	todo：思路不错，最后一步不太理解
	*/
type ListNode struct {
    Val int
    Next *ListNode
    Random *ListNode
}


func copyRandomList(head *ListNode) *ListNode {
	// 复制当前节点
    for p := head; p !=nil;{
		np := &ListNode{}
		// 备份
		next := p.Next
		p.Next = np
		np.Next = next
		// 下一个节点
		p = next
	}
	// 处理random
	for p := head; p != nil; p = p.Next.Next{
		if p.Random != nil{
			p.Next.Random = p.Random.Next
		}
	}
	// 看不懂了
	dummy := &ListNode{}
	// todo：这cur他么是做什么的，看不懂了
	cur := dummy
	for p := head;p != nil;p = p.Next{
		cur.Next = p.Next
		cur = cur.Next
		p = p.Next
	}
	return dummy.Next
}