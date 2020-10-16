/**
判断镜像二叉树问题

递归（dfs）的方式
1.两个根节点的值要相等
2.左边的左子树和右边的右子树对称
3.左边的右子树和右边的左子树对称
*/
func isSymmetric(root *TreeNode) bool {
	// 如果当前root是空节点，返回true
	if root == nil {
		return true
	}
	// 递归左右子节点
	return dfs(root.Left, root.Right)
}

func dfs(left, right *TreeNode) bool {
	// 如果左右子节点都为空，返回true
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	/*
		以下三种情况返回true
		（1）左右子节点value相同
		（2）左子树的左子树 和 右子树的右子树 满足dfs
		（2）左子树的右子树 和 右子树的左子树 满足dfs
	*/
	return left.Val == right.Val && dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}
