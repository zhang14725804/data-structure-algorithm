/*
	给定一个非空二叉树，返回其最大路径和。
	本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点

	向左走，向右走，不走（边的权重给定，不是1）

	思路：枚举所有点
*/

/*
	方法1：递归：自顶向下
	(question)硬是没看懂😅
*/
const min = -(1 << 32) - 1

// 声明最小值
var ans int = min

func maxPathSum(root *TreeNode) int {
	// 本质上是后序遍历(🔥🔥🔥)
	dfs(root)
	return ans
}

func dfs(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}
	// 递归左右节点
	left := dfs(root.Left)
	right := dfs(root.Right)
	// 当前路径和与ans比较
	ans = MaxInt(ans, left+root.Val+right)
	// ps：返回当前节点的最大路径和（如果小于0需要排除）
	return MaxInt(0, root.Val+MaxInt(left, right))
}

