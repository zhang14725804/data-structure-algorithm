package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
101. Symmetric Tree
判断镜像二叉树问题
递归（dfs）的方式
1.两个根节点的值要相等
2.左边的左子树和右边的右子树对称
3.左边的右子树和右边的左子树对称
*/
func isSymmetric01(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

func dfs(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	return left.Val == right.Val && dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}
