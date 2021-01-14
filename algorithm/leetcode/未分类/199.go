/*
	给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

	依次输出二叉树每层最右边的元素。

	思路1：层序遍历，每次取最后一个元素
	思路2：深度优先遍历，从右子树开始
*/
var res []int

func rightSideView(root *TreeNode) []int {
	if root != nil {
		dfs(root, 0)
	}
	return res
}
func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	// 用一个变量记录当前层数，每次保存第一次到达该层的元素。
	if level == len(res) {
		res = append(res, root.Val)
	}
	dfs(root.Right, level+1)
	dfs(root.Left, level+1)
}