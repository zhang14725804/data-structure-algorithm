/*
	方法1：DFS-递归实现
	指针和地址是唯一的难点 😅😅
*/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	dfs(root, &res)
	return res

}
func dfs(node *TreeNode, res *[]int) {
	// 当前节点为空
	if node == nil {
		return
	}
	// 左子树
	dfs(node.Left, res)
	// 指针
	(*res) = append(*res, node.Val)
	// 右子树
	dfs(node.Right, res)
}

/*
	方法1：DFS-递归实现
	用全局变量替代指针地址
*/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 左、跟、右
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return res
}

/*
	思路2：DFS-迭代实现
	question 😅😅😅😅
	(1) 将整棵树的最左边一条链压入栈
	(2) 每次取出栈顶元素，如果它有右子树，将右子树压入栈中
*/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	// 栈
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			// 左，将整棵树的最左边一条链压入栈
			stack = append(stack, root)
			root = root.Left
		} else {
			// 最后一个元素出栈
			sLen := len(stack) - 1
			cur := stack[sLen]
			stack = stack[:sLen]
			// 根
			res = append(res, cur.Val)
			// 右，将整棵树的最右边一条链压入栈
			root = cur.Right
		}
	}
	return res
}

/*
	思路2：DFS-迭代实现
	右根左入栈，出栈刚好是：左根右
	需要使用空节点作为标记 （不好理解）
*/
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		// （1）弹出顶部元素
		cLen := len(stack) - 1
		cnode := stack[cLen]
		stack = stack[:cLen]
		if cnode != nil {
			// （2） 右 😅😅😅
			if cnode.Right != nil {
				stack = append(stack, cnode.Right)
			}
			// 根
			stack = append(stack, cnode)
			// question 😅😅😅😅😅😅 中节点访问过，但是还没有处理，加入空节点做为标记。
			stack = append(stack, nil)
			// 左
			if cnode.Left != nil {
				stack = append(stack, cnode.Left)
			}
		} else {
			// 😅😅😅 只有遇到空节点的时候，才将下一个节点放进结果集
			cLen = len(stack) - 1
			cnode = stack[cLen]
			stack = stack[:cLen]
			res = append(res, cnode.Val)
		}
	}
	return res
}