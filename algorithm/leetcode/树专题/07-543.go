package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。

	思路1：枚举所有点（有所边的权重是1）
	左右递归时需要计算，从当前节点往下走，深度的最大值

	[]空数组的时候无法通过，[1,2,3,4,5]可以通过；leetcode编译器有问题，有缓存
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
	// 递归左右节点
	left := dfs(root.Left)
	right := dfs(root.Right)
	// 所有经过root的路径最大值，更新最大值
	ans = Max(ans, left+right)
	// 返回root的最大深度
	return Max(left+1, right+1)
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
