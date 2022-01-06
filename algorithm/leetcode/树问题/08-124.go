/*
	给定一个非空二叉树，返回其最大路径和。
	本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点

	向左走，向右走，不走（边的权重给定，不是1）
	leetcode 无法通过
	0106 懵逼
*/

/*
	方法1：递归：自顶向下
*/
var res int

func maxPathSum(root *TreeNode) int {
	//
	res = -(1 << 32)
	dfs(root)
	return res
}

// dfs 返回当前节点最大路径和
func dfs(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}
	// 递归左右节点
	left := MaxInt(dfs(root.Left), 0)
	right := MaxInt(dfs(root.Right), 0)
	// 😅😅😅 更新ans值
	res = MaxInt(res, root.Val+left+right)
	// 😅😅😅  三种情况：向左走，向右走，不走
	return root.Val + MaxInt(left, right)
}
