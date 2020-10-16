/*
	二叉树的后序遍历（ps：前序中序后序说的是根元素位置😂）
	迭代和递归（todo）
*/
func postorderTraversal1(root *TreeNode) []int {
	// 居然连递归方法都写错了
	res := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)
	return res
}

func postorderTraversal(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode

	for root != nil || len(stk) > 0 {
		// 遍历左子树
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		// 去除栈顶元素
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		// 判断是否存在右子树；root.Right == prev 什么意思（question）
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			// 遍历右子树
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}