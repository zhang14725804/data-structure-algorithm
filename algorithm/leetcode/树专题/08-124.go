package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	给定一个非空二叉树，返回其最大路径和。
	本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点

	向左走，向右走，不走（边的权重给定，不是1）

	思路：枚举所有点

	[0]不通过，[1,2,3]可以通过；leetcode编译器有问题，有缓存
*/
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// 声明最小值
var ans int = MinInt

func maxPathSum(root *TreeNode) int {
	dfs(root)
	return ans
}

// 返回从root往下递归的最大值
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归左右节点
	left := dfs(root.Left)
	right := dfs(root.Right)
	// 当前路径和与ans比较
	ans = Max(ans, left+root.Val+right)
	// ps：返回当前节点的最大路径和（如果小于0需要排除）
	return Max(0, root.Val+Max(left, right))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
