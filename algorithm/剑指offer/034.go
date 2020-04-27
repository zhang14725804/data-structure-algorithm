/*
	链表中环的入口结点（单链表中只有一个环😅？好像是这样的）
	todo：经典问题
	首先，first走一步，second走两步，第一次相遇之后，first回到起点，然后first和second都开始走，每次都走一步，再次相遇的点就是入口
	todo：如何证明，如何证明才是最重要的
*/
type ListNode struct {
    Val int
    Next *ListNode
}
func entryNodeOfLoop(head *ListNode)  *ListNode{
	// i,j 快慢指针
	i,j := head,head
	for i!=nil && j!=nil{
		i = i.Next
		j = j.Next
		// j 每次走两步
		if j!=nil{
			j = j.Next
		}
		// 快慢指针第一次相遇
		if i==j{
			// 满指针回到起点
			i = head
			for i != j{
				i = i.Next
				j = j.Next
			}
			return i
		}
	}
	return nil
}