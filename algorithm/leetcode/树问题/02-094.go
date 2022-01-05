/*
	二叉树中序序遍历
	0105
*/

/*
	方法1：DFS-递归实现
	指针、地址 😅😅
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
	// TODOS：：这个什么语法
	(*res) = append(*res, node.Val)
	// 右子树
	dfs(node.Right, res)
}

/*
	方法1：DFS-递归实现
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
	stack := make([]*TreeNode, 0)
	if root == nil {
		return res
	}
	cur := root
	// 迭代条件 😅😅😅
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			// （1）将整棵树的最左边一条链压入栈
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			// （2）从栈里弹出最后一个元素
			cLen := len(stack) - 1
			cur = stack[cLen]
			stack = stack[:cLen]
			// （3）根、右
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

/*
	思路2：DFS-迭代实现，使用空节点作为标记
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
			// （2） 😅😅😅
			// 右
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
			// 只有遇到空节点的时候，才将下一个节点放进结果集
			cLen = len(stack) - 1
			cnode = stack[cLen]
			stack = stack[:cLen]
			res = append(res, cnode.Val)
		}
	}
	return res
}