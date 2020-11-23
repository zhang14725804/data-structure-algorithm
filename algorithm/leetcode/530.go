/*
	给你一棵所有节点为非负值的 **二叉搜索树**，请你计算树中任意两节点的差的绝对值的最小值。
	二叉搜索树性值：二叉搜索树中序遍历得到的值序列是递增有序的
	方法：中序遍历即可
*/
var ans int = 1 << 32
var last int
var isFirst bool = true

func getMinimumDifference(root *TreeNode) int {
	dfs(root)
	return ans
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if isFirst {
		isFirst = false
	} else {
		ans = MaxMin(ans, root.Val-last)
	}
	last = root.Val
	dfs(root.Right)
}