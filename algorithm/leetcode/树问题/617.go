/*
	方法1：DFS-递归
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// base case
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 😅😅 我居然还新建了root节点，然后再向root节点添加左右子树，之后再返回root节点
	// 前序遍历：根左右
	root1.Val = root1.Val + root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

/*
	方法2：BFS-层序遍历
*/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// base case
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	var queue []*TreeNode
	queue = append(queue, root1)
	queue = append(queue, root2)

	for len(queue) > 0 {
		cnode1 := queue[0]
		queue = queue[1:]
		cnode2 := queue[0]
		queue = queue[1:]
		// 合并两个节点
		cnode1.Val = cnode1.Val + cnode2.Val
		// 两棵树左节点都不为空
		if cnode1.Left != nil && cnode2.Left != nil {
			queue = append(queue, cnode1.Left)
			queue = append(queue, cnode2.Left)
		}
		// 两棵树右节点都不为空
		if cnode1.Right != nil && cnode2.Right != nil {
			queue = append(queue, cnode1.Right)
			queue = append(queue, cnode2.Right)
		}
		// 😅😅 没想到这一步。左右节点不同时为空的情况
		if cnode1.Left == nil && cnode2.Left != nil {
			cnode1.Left = cnode2.Left
		}
		if cnode1.Right == nil && cnode2.Right != nil {
			cnode1.Right = cnode2.Right
		}
	}
	return root1
}