package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
/**
	98. Validate Binary Search Tree

	自上而下递归判断
	
	https://www.cnblogs.com/coffeeSS/p/5452719.html
*/
func isValidBST(root *TreeNode) bool {
	const MaxUint = ^uint(0) 
	const MinUint = 0 
	const MaxInt = int(MaxUint >> 1) 
	const MinInt = -MaxInt - 1
    return dfs(root,MinInt,MaxInt)
}

func dfs(root *TreeNode,minv,maxv int) bool{
	if root==nil{
		return true
	}
	if root.Val < minv || root.Val>maxv{
		return false
	}

	return dfs(root.Left,minv,root.Val-1) && dfs(root.Right,root.Val+1,maxv)
}