package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
	101. Symmetric Tree
	判断镜像二叉树问题

	迭代的方式（不同的遍历方式，然后比较节点即可）
	左边：左中右
	右边：右中左
*/
func isSymmetric(root *TreeNode) bool {
    if root==nil{
		return true
	}
	return dfs(root.Left,root.Right)
}

func dfs(left,right *TreeNode) bool{
	if left==nil || right==nil{
		return left==nil && right==nil
	}
	return left.Val == right.Val && dfs(left.Left,right.Right) && dfs(left.Right,right.Left)
}

