package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	124. Binary Tree Maximum Path Sum

	向左走，向右走，不走（和543题类似）

	todos:依旧无法通过临界条件[0]不通过，[1,2,3]可以通过
*/
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// 声明最小值
var ans int = MinInt
func maxPathSum(root *TreeNode) int {
	dfs(root)
	return ans
}

func dfs(root *TreeNode) int{
	if root ==nil{
		return 0
	}
	left:=dfs(root.Left)
	right:=dfs(root.Right)
	ans = Max(ans,left+root.Val+right)
	return Max(0,root.Val + Max(left,right))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}