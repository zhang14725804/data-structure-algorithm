/*
	输入一棵二叉树的根结点，判断该树是不是平衡二叉树
	如果某二叉树中任意结点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树
*/
var ans = true

func isBalanced(root *TreeNode) bool {
	dfs(root)
	return ans
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	if abs(left, right) > 1 {
		ans = false
	}
	return max(left, right) + 1
}

// 工具函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}