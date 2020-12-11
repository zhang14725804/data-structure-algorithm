/*
	给定一个带有头结点 head 的非空单链表，返回链表的中间结点。
	如果有两个中间结点，则返回第二个中间结点。

	思路：快慢指针；两个指针，一个每次走一步，一个每次走两步。
*/
func middleNode(head *ListNode) *ListNode {
    slow := head
    fast := head
    for fast.Next != nil && fast.Next.Next != nil{
        fast = fast.Next.Next
        slow = slow.Next 
	}
	// 偶数长度
    if fast.Next != nil{
        return slow.Next
	}
	// 奇数长度
    return slow
}