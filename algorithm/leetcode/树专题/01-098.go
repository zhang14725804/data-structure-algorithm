package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	给定一个二叉树，判断其是否是一个有效的二叉搜索树。

	假设一个二叉搜索树具有如下特征：

	节点的左子树只包含小于当前节点的数。
	节点的右子树只包含大于当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。


	两种方式：自顶向下、自底向上
*/

/*
	思路：自顶向下
	左右子树的值都要在一个区间当中
*/
func isValidBST1(root *TreeNode) bool {
	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
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
	自底向上（😅，这种做法代码有问题）

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
