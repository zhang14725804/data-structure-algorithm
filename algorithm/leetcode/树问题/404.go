
/*
	方法1：DFS-递归
	前序遍历和后序遍历都可以
*/
func sumOfLeftLeaves(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}
	res := 0
	// 😅 通过节点的父节点来判断其左孩子是不是左叶子
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}
	return res + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

/*
	方法2：DFS-迭代法
	用【栈】先进后出模拟深度优先遍历
*/
func sumOfLeftLeaves(root *TreeNode) int {
	var stack []*TreeNode
	var res int
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		// 后进先出
		cLen := len(stack) - 1
		cnode := stack[cLen]
		stack = stack[:cLen]

		if cnode.Left != nil && cnode.Left.Left == nil && cnode.Left.Right == nil {
			res += cnode.Left.Val
		}
		// 先右后左
		if cnode.Right != nil {
			stack = append(stack, cnode.Right)
		}
		if cnode.Left != nil {
			stack = append(stack, cnode.Left)
		}
	}
	return res
}