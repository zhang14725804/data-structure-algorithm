/*
	思路1：递归
	1.
*/
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 左右子树的深度
		left := dfs(root.Left)
		right := dfs(root.Right)
		// 计算经过root的路径最大值，更新最大值
		ans = Max(ans, left+right)
		// 返回节点最大深度
		return Max(left, right) + 1
	}
	dfs(root)
	return ans
}

