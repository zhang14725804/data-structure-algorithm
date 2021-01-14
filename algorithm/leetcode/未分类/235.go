/*
	给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

	二叉搜索树的定义：

		若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
		若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
		任意节点的左、右子树也分别为二叉查找树；
		没有键值相等的节点。
*/

/*
	方法1：递归

	如果给定的两个节点的值都小于根节点的值，那么最近的共同祖先一定在左子树
	如果给定的两个节点的值都大于根节点的值，那么最近的共同祖先一定在右子树
	如果一个大于等于、一个小于等于根节点的值，那么当前根节点就是最近的共同祖先了
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 保持p.val < q.val，省去复杂的判断（妙啊😄）
	if p.Val > q.Val {
		return lowestCommonAncestor(root, q, p)
	}
	if p.Val == root.Val || q.Val == root.Val {
		return root
	}

	if q.Val < root.Val { // 较大的值小于root
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val { // 较小的值大于root
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

// 方法2：迭代
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pVal, qVal := p.Val, q.Val
	if pVal == root.Val || qVal == root.Val {
		return root
	}
	if pVal > qVal {
		pVal, qVal = qVal, pVal
	}
	// 最后如果不使用else，直接使用return，二者有什么区别（todo）
	for {
		if qVal < root.Val {
			root = root.Left
		} else if pVal > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
}