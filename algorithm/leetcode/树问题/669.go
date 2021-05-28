/*
	给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。
	通过修剪二叉搜索树，使得所有节点的值在[low, high]中。
	修剪树不应该改变保留在树中的元素的相对结构（即，如果没有被移除，原有的父代子代关系都应当保留）。 可以证明，存在唯一的答案。

	所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。
*/

/*
	方法1：DFS-递归
	😅😅😅 多余的节点究竟是如何从二叉树中移除的呢？
*/
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// base case
	if root == nil {
		return root
	}
	// 😅😅😅 如果root（当前节点）的元素小于low的数值，那么应该递归右子树，并返回右子树符合条件的头结点（删除了左子树）
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	// 😅😅😅 如果root(当前节点)的元素大于high的，那么应该递归左子树，并返回左子树符合条件的头结点。（删除了右子树）
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	// 😅😅😅 将下一层处理完左子树的结果赋给root->left，处理完右子树的结果赋给root->right
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

/*
	方法2：DFS-迭代法 question 没理解😅
	🔥🔥🔥 因为二叉搜索树的有序性，不需要使用栈模拟递归的过程。
*/
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}
	// 😅😅😅 处理头结点，让root移动到[L, R] 范围内，注意是左闭右闭
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	cnode := root
	// 😅😅😅 此时root已经在[L, R] 范围内，处理左孩子元素小于L的情况
	for cnode != nil {
		for cnode.Left != nil && cnode.Left.Val < low {
			cnode.Left = cnode.Left.Right
		}
		cnode = cnode.Left
	}
	cnode = root
	// 😅😅😅 此时root已经在[L, R] 范围内，处理右孩子大于R的情况
	for cnode != nil {
		for cnode.Right != nil && cnode.Right.Val > high {
			cnode.Right = cnode.Right.Left
		}
		cnode = cnode.Right
	}
	return root
}