/*
	😅😅😅 一顿操作猛如虎
	(question)问题是，对于某一个节点 root，他只能管得了自己的左右子节点，怎么把 root 的约束传递给左右子树呢？
	question 😅😅 [5,4,6,null,null,3,7] 测试不通过

	陷阱1：不能单纯的比较左节点小于中间节点，右节点大于中间节点就完事了。
	陷阱2：样例中最小节点 可能是int的最小值，如果这样使用最小的int来比较也是不行的。
*/
func isValidBST0(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Val <= root.Left.Val {
		return false
	}
	if root.Right != nil && root.Val >= root.Right.Val {
		return false
	}
	return isValidBST0(root.Left) && isValidBST0(root.Right)
}

/*
	方法1：DFS-递归
	前序遍历：根左右
	根据 BST 的定义，root 的整个左子树都要小于 root.val，整个右子树都要大于 root.val
	😅😅😅 通过使用辅助函数，增加函数参数列表，在参数中携带额外信息，将这种约束传递给子树的所有节点
*/
func isValidBST(root *TreeNode) bool {
	return dfs(root, nil, nil)
}

func dfs(root, min, max *TreeNode) bool {
	// base case
	if root == nil {
		return true
	}
	// 整个右子树都要大于 root.val
	if min != nil && root.Val <= min.Val {
		return false
	}
	// 整个左子树都要小于 root.val
	if max != nil && root.Val >= max.Val {
		return false
	}
	// 递归执行，并传递【min、max约束条件】😅
	return dfs(root.Left, min, root) && dfs(root.Right, root, max)
}

/*
	DFS-递归 😄😄😄 机智
	***************左右子树的值都要在一个区间当中********************
*/
func isValidBST(root *TreeNode) bool {
	min := -(1 << 32)
	return dfs(root, min, -min)
}
func dfs(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val < min || root.Val > max {
		return false
	}
	return dfs(root.Left, min, root.Val-1) && dfs(root.Right, root.Val+1, max)
}

/*
	方法3：DFS-递归
	利用二叉搜索树性质，【中序遍历（左根右）】是排序数组
	二叉树中序遍历，判断遍历结果是否【有序】，是否有重复元素
	只需要进行一次中序遍历，将遍历结果保存，然后判断该数组是否是从小到大排列的即可。
	由于我们只需要临近的两个数的相对关系，所以我们只需要在遍历过程中，把当前遍历的结果和上一个结果比较即可。
*/
var prev *TreeNode

func isValidBST(root *TreeNode) bool {
	return dfs(root)
}
func dfs(root *TreeNode) bool {
	// base case
	if root == nil {
		return true
	}
	// 左
	lv := dfs(root.Left)
	// 根。由于我们只需要临近的两个数的相对关系，所以我们只需要在遍历过程中，把当前遍历的结果和上一个结果比较即可
	if prev != nil && root.Val <= prev.Val {
		return false
	}
	// 跟新前一个节点
	prev = root
	// 右
	rv := dfs(root.Right)
	return lv && rv
}

/*
	方法4：DFS-迭代法
	利用二叉搜索树性质，【中序遍历】是排序数组
*/
func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var stack []*TreeNode
	cnode := root
	// 为什么不直接把root 放入stack
	for cnode != nil || len(stack) > 0 {
		if cnode != nil {
			// 入栈 😅
			stack = append(stack, cnode)
			// 左
			cnode = cnode.Left
		} else {
			cLen := len(stack) - 1
			cnode = stack[cLen]
			stack = stack[:cLen]
			// 😅 由于我们只需要临近的两个数的相对关系，所以我们只需要在遍历过程中，把当前遍历的结果和上一个结果比较即可
			if prev != nil && cnode.Val <= prev.Val {
				return false
			}
			// 跟新prev节点
			prev = cnode
			// 右
			cnode = cnode.Right
		}
	}
	return true
}