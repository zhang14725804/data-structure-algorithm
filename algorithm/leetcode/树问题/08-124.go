/*
	方法1：递归：自顶向下
	不懂 😅😅😅
*/
var res int = -(1 << 32)

func maxPathSum(root *TreeNode) int {
	dfs(root)
	return res
}

// dfs 返回当前节点最大路径和
func dfs(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}

	// 递归计算左右子节点的最大贡献值
	// 只有在最大贡献值大于 0 时，才会选取对应子节点
	left := Max(dfs(root.Left), 0)
	right := Max(dfs(root.Right), 0)
	// 😅😅😅 更新ans值，节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
	res = Max(res, root.Val+left+right)
	// 😅😅😅  三种情况：向左走，向右走，不走。返回节点的最大贡献值
	return root.Val + Max(left, right)
}
