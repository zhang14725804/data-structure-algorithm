/*
	给定一个二叉树，找出其最小深度。

	最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

	说明: 叶子节点是指没有子节点的节点。

	时间复杂度O(n)
	ps：自下而上的方式
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 计算左右子树
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	// 左边或者右边是空
	if left == 0 || right == 0 {
		return left + right + 1
	}
	// 左右子树都不为空
	return min(left, right) + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}