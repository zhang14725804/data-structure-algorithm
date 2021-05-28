/*
	给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

	二叉搜索树的定义：

		若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
		若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
		任意节点的左、右子树也分别为二叉查找树；
		没有键值相等的节点。
*/

/*
	方法1：DFS-递归

	在有序树里，如果判断一个节点的左子树里有p，右子树里有q呢？
	🔥🔥🔥其实只要从上到下遍历的时候，cur节点是数值在[p, q]区间中则说明该节点cur就是最近公共祖先了。

	如果给定的两个节点的值都小于根节点的值，那么最近的共同祖先一定在左子树
	如果给定的两个节点的值都大于根节点的值，那么最近的共同祖先一定在右子树
	如果一个大于等于、一个小于等于根节点的值，那么当前根节点就是最近的共同祖先了
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

/*
	方法2：DFS-迭代
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}