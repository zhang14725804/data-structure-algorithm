/*
	翻转一棵二叉树。
*/

/*
	方法1：dfs-深度优先遍历，递归实现（自顶向下）
	前序遍历和后序遍历都可以 😅
*/
func invertTree(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 处理当前节点
		root.Left, root.Right = root.Right, root.Left
		// 递归左右子树
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return root
}

/*
	方法2：BFS-广度优先遍历，借助【队列】实现层序遍历
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		// 从队首获取元素，同时获取的这个元素将从原队列删除；
		cnode := queue[0]
		queue = queue[1:]
		// 交换左右子树
		cnode.Left, cnode.Right = cnode.Right, cnode.Left
		// 递归左右子树
		// 【队列】，先进先出，所以先push左子树，后push右子树
		if cnode.Left != nil {
			queue = append(queue, cnode.Left)
		}
		if cnode.Right != nil {
			queue = append(queue, cnode.Right)
		}
	}
	return root
}

/*
	方法3：DFS-深度优先遍历，迭代法实现
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		// 从队首获取元素，同时获取的这个元素将从原队列删除；
		cLen := len(stack) - 1
		cnode := stack[cLen]
		stack = stack[:cLen]
		// 交换左右子树
		cnode.Left, cnode.Right = cnode.Right, cnode.Left
		// 递归左右子树
		// 【栈】，后进先出，所以先push右子树，后push左子树
		if cnode.Right != nil {
			stack = append(stack, cnode.Right)
		}
		if cnode.Left != nil {
			stack = append(stack, cnode.Left)
		}
	}
	return root
}