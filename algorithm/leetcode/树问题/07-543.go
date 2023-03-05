/*
	思路1：枚举所有点（有所边的权重是1）
	左右递归时需要计算，从当前节点往下走，深度的最大值
	😅😅😅 没有思路
*/
var ans int

func diameterOfBinaryTree(root *TreeNode) int {
	dfs(root)
	return ans
}

func dfs(root *TreeNode) int {
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
