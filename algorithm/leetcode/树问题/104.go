/*
	给定一个二叉树，找出其最大深度。

	二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

	说明: 叶子节点是指没有子节点的节点。
*/

// DFS递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return MaxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// BFS 层序遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	level := 0

	for len(stack) > 0 {
		levelNum := len(stack)
		for i := 0; i < levelNum; i++ {
			/*
				careful： 从队尾操作是错的😅，要从队头开始操作（ 为什么呢 question ）
				current := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			*/
			current := stack[0]
			stack = stack[1:]
			if current.Left != nil {
				stack = append(stack, current.Left)
			}
			if current.Right != nil {
				stack = append(stack, current.Right)
			}
		}
		level++
	}
	return level
}