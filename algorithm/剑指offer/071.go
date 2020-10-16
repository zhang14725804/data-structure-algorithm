/*
	输入一棵二叉树的根结点，求该树的深度(早上在路上想到这个问题了😃)
	根节点到叶子节点当中最长的路径

	左子树的深度和右子树的最大值，再加上1（递归）
*/
func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(treeDepth(root.Left), treeDepth(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}