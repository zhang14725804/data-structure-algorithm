/*
	方法1：DFS-递归
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return dfs(p, q)
}

// 后序遍历：根左右
func dfs(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	// 根
	if p.Val != q.Val {
		return false
	}
	// 左右
	return dfs(p.Left, q.Left) && dfs(p.Right, q.Right)
}