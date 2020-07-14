/*
	给定一个二叉树，找出其最大深度。

	二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

	说明: 叶子节点是指没有子节点的节点。
*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	// dfs即可
	return compare(maxDepth(root.Left),maxDepth(root.Right),true)+1
}

func compare(a, b int, max bool) int {
	// max 是否返回最大值
	if a > b {
		if max == true {
			return a
		}
		return b
	}
	if max == true {
		return b
	}
	return a
}