/*
	给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
	本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

	思路：找中间结点是难点
*/

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	// 链表长度
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	if n == 1 {
		return &TreeNode{Val: head.Val}
	}
	// 找到中间节点
	cur := head
	for i := 0; i < n/2-1; i++ {
		cur = cur.Next
	}
	root := &TreeNode{Val: cur.Next.Val}
	// 递归建立左右子节点
	root.Right = sortedListToBST(cur.Next.Next)
	// 中间节点置为nil
	cur.Next = nil
	root.Left = sortedListToBST(head)
	return root
}