/*
	判断一棵二叉树是不是对称的
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}
func dfs(l, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == nil && r == nil
	}
	if l.Val != r.Val {
		return false
	}
	return dfs(l.Left, r.Right) && dfs(l.Right, r.Left)
}