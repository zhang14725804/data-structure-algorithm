/*
	方法1：DFS-递归实现
*/
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		// base case
		if root == nil {
			return
		}
		// 左根右
		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return ans
}

/*
	思路2：用栈实现 😅😅😅
	1. 将整棵树的最【左】边一条链压入栈
	2. 取出栈顶【根】元素加入集合
	3. 如果它有【右】子树，将右子树压入栈中
*/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	// 这里是for 😅
	for root != nil || len(stack) > 0 {
		if root != nil {
			// 😅 左，将整棵树的最左边一条链压入栈
			stack = append(stack, root)
			root = root.Left
		} else {
			// 最后一个元素出栈
			sLen := len(stack) - 1
			cur := stack[sLen]
			stack = stack[:sLen]
			// 根
			res = append(res, cur.Val)
			// 😅 右，将整棵树的最右边一条链压入栈
			root = cur.Right
		}
	}
	return res
}

/*
	思路2：用层序遍历实现 😅😅😅😅
	1. 右根左入栈（出栈刚好是左根右）
	2. 使用空节点作为标记 （难以理解）
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