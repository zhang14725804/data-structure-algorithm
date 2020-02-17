package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	236. Lowest Common Ancestor of a Binary Tree（公共祖先）

	// 如果以root为根的子树中包含p和q，则返回他们的最近公共祖先
	// 如果只包含p，返回p
	// 如果只包含q，返回q
	// 如果都不包含，则返回null
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
