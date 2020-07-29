/*
	给定一个二叉树和一个目标和，
	判断该树中是否存在根节点到叶子节点的路径，
	（ps：那走到一半，并没有走到叶子节点，但是sum符合，怎么办）
	这条路径上所有节点值相加等于目标和。

	从上而下考虑
	f(u)：从根节点走到u的权值之和

*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	sum-=root.Val
	if root.Left == nil && root.Right == nil{
		return sum == 0
	}
	return root.Left!=nil && hasPathSum(root.Left,sum) || root.Right!=nil && hasPathSum(root.Right,sum)
}