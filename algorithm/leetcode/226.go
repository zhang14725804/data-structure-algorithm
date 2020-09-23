/*
	翻转一棵二叉树。
*/
// 方法1：递归dfs
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 方法2：迭代bfs，todo
