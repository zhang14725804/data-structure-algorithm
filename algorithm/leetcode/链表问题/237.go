// Leetcode237 Delete Node in a Linked List
func Leetcode237(node *ListNode) {
	/*
		让下个节点的val覆盖当前节点的val，然后删除下个节点，让本届点的next指向下个节点的下个节点
	*/
	// 方法1
	// node.Val = node.Next.Val
	// node.Next = node.Next.Next
	// 方法2
	*node = *(node.Next)
}
