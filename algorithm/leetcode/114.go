/*
	给定一个二叉树，原地将它展开为一个单链表。

	(1)如果存在左子树
	(2)否则，遍历右子树

	todo：有点绕
*/

func flatten(root *TreeNode) {
	for root != nil {
		p := root.Left
		if p != nil {
			for p.Right != nil {
				p = p.Right
			}
			p.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		root = root.Right
	}
}