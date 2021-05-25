/*
	二叉树的后序遍历
*/

/*
	方法1：DFS-递归实现
*/
func postorderTraversal1(root *TreeNode) []int {
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

/*
	方法2：迭代实现（ question 妙啊 😅😅）
	先序遍历（根左右）-> 根右左 -> 反转数组->左右根
*/
func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack = append(stack, root)
	// 【根右左】的遍历结果
	for len(stack) > 0 {
		sLen := len(stack)
		for i := 0; i < sLen; i++ {
			// 栈，后进先出
			cLen := len(stack) - 1
			cur := stack[cLen]
			stack = stack[:cLen]
			res = append(res, cur.Val)
			// 左子树后入栈
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			// 右子树先入栈
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
		}
	}
	// 反转数组
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

/*
	方法2：迭代实现，使用空节点作为标记（TODO 😅）
*/
func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		// 弹出顶部元素
		cLen := len(stack) - 1
		cnode := stack[cLen]
		stack = stack[:cLen]
		if cnode != nil {
			// 根
			stack = append(stack, cnode)
			// question 😅😅😅😅😅😅 中节点访问过，但是还没有处理，加入空节点做为标记。
			stack = append(stack, nil)
			// 右
			if cnode.Right != nil {
				stack = append(stack, cnode.Right)
			}
			// 左
			if cnode.Left != nil {
				stack = append(stack, cnode.Left)
			}
		} else {
			// 只有遇到空节点的时候，才将下一个节点放进结果集
			cLen = len(stack) - 1
			cnode = stack[cLen]
			stack = stack[:cLen]
			res = append(res, cnode.Val)
		}
	}
	return res
}

