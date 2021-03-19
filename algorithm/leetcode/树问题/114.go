/*
	给定一个二叉树，原地将它展开为一个单链表。
*/

/*
	方法1：(question)
	(1)如果存在左子树,遍历左子树
	(2)否则，遍历右子树
*/
func flatten1(root *TreeNode) {
	for root != nil {
		p := root.Left
		// 遍历左子树
		if p != nil {
			// 遍历右子树
			for p.Right != nil {
				p = p.Right
			}
			// 将本来的右子树连接到当前右子树末端
			p.Right = root.Right
			// 将左子树作为右子树
			root.Right = root.Left
			root.Left = nil
		}
		// 遍历右子树
		root = root.Right
	}
}

/*
	方法2：递归(question)
*/
func flatten(root *TreeNode) {
	// base case
	if root == nil {
		return
	}
	// 后序遍历（左右根）
	flatten(root.Left)
	flatten(root.Right)
	// 左右子树已经被拉成一条链表
	left := root.Left
	right := root.Right
	// 将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 将本来的右子树连接到当前右子树末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}