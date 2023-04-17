/*
	方法1：DFS-递归实现
*/
func preorderTraversal(root *TreeNode) []int {
	var res []int
	// 确定递归函数的参数和返回值
	var dfs func(root *TreeNode)
	// 根左右
	dfs = func(root *TreeNode) {
		// base case
		if root == nil {
			return
		}
		// 根左右
		res = append(res, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

/*
	用栈模拟递归
	😅😅😅
*/
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	stk := []*TreeNode{}
	for root != nil || len(stk) > 0 {
		if root != nil {
			res = append(res, root.Val)
			// 这里不理解 😅，不应该是先root = root.Left，后入栈么
			stk = append(stk, root)
			root = root.Left
		} else {
			root = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			root = root.Right
		}
	}
	return res
}

/*
	方法2：迭代【栈】实现
	根右左的方式入栈😅😅😅
*/
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	stk := []*TreeNode{}
	if root != nil {
		stk = append(stk, root)
	}
	for len(stk) > 0 {
		for i := 0; i < len(stk); i++ {
			node := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			//
			res = append(res, node.Val)
			//  入栈先右后左！！ 😅😅
			if node.Right != nil {
				stk = append(stk, node.Right)
			}
			if node.Left != nil {
				stk = append(stk, node.Left)
			}
		}
	}
	return res
}

/*
	迭代【栈】实现
	使用【😅空节点😅】作为标记
	TODO 😅😅😅
*/
func preorderTraversal(root *TreeNode) []int {
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
		// 右左根【空节点】的顺序入栈
		if cnode != nil {
			if cnode.Right != nil {
				stack = append(stack, cnode.Right)
			}
			if cnode.Left != nil {
				stack = append(stack, cnode.Left)
			}
			stack = append(stack, cnode)
			// question 😅😅😅😅😅😅 中节点访问过，但是还没有处理，加入空节点做为标记。
			stack = append(stack, nil)
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

