/*
	给定一个二叉树，判断它是否是高度平衡的二叉树。

	本题中，一棵高度平衡二叉树定义为：

	一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
*/
var ans bool

func isBalanced(root *TreeNode) bool {
	ans = true
	dfs(root)
	return ans
}

// 返回当前结点高度
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归左右子节点
	lh := dfs(root.Left)
	rh := dfs(root.Right)
	// 判断是否是平衡二叉树
	if abs(lh, rh) > 1 {
		ans = false
	}
	return compare(lh, rh, true) + 1
}
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}