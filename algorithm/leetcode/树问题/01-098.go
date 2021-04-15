/*
	给定一个二叉树，判断其是否是一个有效的二叉搜索树。

	假设一个二叉搜索树具有如下特征：
		节点的左子树只包含小于当前节点的数。
		节点的右子树只包含大于当前节点的数。
		所有左子树和右子树自身必须也是二叉搜索树。
*/

/*
	刚开始的想法，直接一个递归，看起来是对的，其实错了
	[5,4,6,null,null,3,7] 不通过
	根据 BST 的定义，root 的整个左子树都要小于 root.val，整个右子树都要大于 root.val
	(question)问题是，对于某一个节点 root，他只能管得了自己的左右子节点，怎么把 root 的约束传递给左右子树呢？
	question 😅😅 [5,4,6,null,null,3,7] 测试不通过，逻辑毫无破绽
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
	根据 BST 的定义，root 的整个左子树都要小于 root.val，整个右子树都要大于 root.val
	通过使用辅助函数，增加函数参数列表，在参数中携带额外信息，将这种约束传递给子树的所有节点
*/
func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}
func helper(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}

	return helper(root.Left, min, root) && helper(root.Right, root, max)
}

/*
	方法1：自顶向下
	***************左右子树的值都要在一个区间当中********************
*/
func isValidBST1(root *TreeNode) bool {
	const MaxInt = (1 << 32)
	const MinInt = -MaxInt - 1
	return dfs1(root, MinInt, MaxInt)
}

func dfs1(root *TreeNode, minv, maxv int) bool {
	if root == nil {
		return true
	}
	if root.Val < minv || root.Val > maxv {
		return false
	}
	return dfs1(root.Left, minv, root.Val-1) && dfs1(root.Right, root.Val+1, maxv)
}

/*
	方法2：自底向上(question 😅，这种做法代码有问题)
	todo
	分别求左右子树最大值和最小值，根节点的值大于左子树最大值，小于右子树最小值
*/
func isValidBST2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var maxv int
	var minv int
	return dfs2(root, maxv, minv)
}

func dfs2(root *TreeNode, maxv, minv int) bool {
	maxv = root.Val
	minv = root.Val
	if root.Left != nil {
		var nowMax, nowMin int
		if dfs2(root.Left, nowMax, nowMin) == false {
			return false
		}
		if nowMax >= root.Val {
			return false
		}
		maxv = compare(maxv, nowMax, true)
		minv = compare(minv, nowMin, false)
	}
	if root.Right != nil {
		var nowMax, nowMin int
		if dfs2(root.Right, nowMax, nowMin) == false {
			return false
		}
		if nowMin <= root.Val {
			return false
		}
		maxv = compare(maxv, nowMax, true)
		minv = compare(minv, nowMin, false)
	}
	return true
}
