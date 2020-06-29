package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	543. Diameter of Binary Tree(树的直径----树当中最长的路径（左右两个方向）)

	枚举所有最高点

	左右递归时需要计算，从当前节点往下走，深度的最大值

	todos：：[]空数组的时候无法通过，[1,2,3,4,5]可以通过
*/
var ans int

func diameterOfBinaryTree(root *TreeNode) int {
	dfs(root)
	return ans
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)

	ans = Max(ans, left+right)
	return Max(left+1, right+1)
}
