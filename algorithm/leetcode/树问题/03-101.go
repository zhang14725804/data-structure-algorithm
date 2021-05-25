/*
	判断镜像二叉树问题
*/

/*
	方法1：DFS-递归实现

	1.两个根节点的值要相等
	2.左边的左子树和右边的右子树对称
	3.左边的右子树和右边的左子树对称

	判断两个子树的内侧节点和外侧节点是否相等：
	左子树遍历顺序😅😅😅：左右根
	右子树遍历顺序😅😅😅：右左根
*/
func isSymmetric(root *TreeNode) bool {
	var dfs func(left, right *TreeNode) bool
	dfs = func(left, right *TreeNode) bool {
		// 如果左右子节点都为空，返回true，其他情况为false
		if left == nil || right == nil {
			return left == nil && right == nil
		}
		/*
			以下三种情况都为true，返回true 😅😅😅
			（1）左右子节点value相同
			（2）左子树的左子树 和 右子树的右子树 满足dfs
			（2）左子树的右子树 和 右子树的左子树 满足dfs
		*/
		return left.Val == right.Val && dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	// 如果当前root是空节点，返回true
	if root == nil {
		return true
	}
	// 递归左右子节点
	return dfs(root.Left, root.Right)
}

/*
	方法2：使用【队列】实现
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue []*TreeNode
	// 左右子树加入队列
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)
	for len(queue) > 0 {
		// 先进先出。从队列中取出左右子树，判断
		lNode := queue[0]
		queue = queue[1:]
		rNode := queue[0]
		queue = queue[1:]
		// 😅 左右子树都为空，符合条件
		if lNode == nil && rNode == nil {
			continue
		}
		// 😅 左右子树不同是为空，或者同时不为空但val不相等
		if lNode == nil || rNode == nil || lNode.Val != rNode.Val {
			return false
		}
		// 😅 按顺序将左右子树的左右子树加入队列
		queue = append(queue, lNode.Left)
		queue = append(queue, rNode.Right)
		queue = append(queue, lNode.Right)
		queue = append(queue, rNode.Left)
	}
	return true
}

/*
	方法3：使用【栈】实现
	直接将队列改成栈，即可 😅😅
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var stack []*TreeNode
	stack = append(stack, root.Left)
	stack = append(stack, root.Right)
	var cLen int
	for len(stack) > 0 {
		// 后进先出
		cLen = len(stack) - 1
		lNode := stack[cLen]
		stack = stack[:cLen]
		cLen = len(stack) - 1
		rNode := stack[cLen]
		stack = stack[:cLen]
		if lNode == nil && rNode == nil {
			continue
		}
		if lNode == nil || rNode == nil || lNode.Val != rNode.Val {
			return false
		}
		stack = append(stack, lNode.Left)
		stack = append(stack, rNode.Right)
		stack = append(stack, lNode.Right)
		stack = append(stack, rNode.Left)
	}
	return true
}
