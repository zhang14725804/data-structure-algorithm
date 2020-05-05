/*
	输入一棵二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所有路径
	todo：代码有问题，测试不通过
*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var ans [][]int
var path []int

func findPath(root *TreeNode, sum int) [][]int {
	dfs(root, sum)
	return ans
}

func dfs(root *TreeNode, sum int){
	// 空节点
	if root == nil{
		return
	}
	path = append(path, root.Val)
	sum -= root.Val
	// 叶子节点，而且满足要求
	if root.Left == nil && root.Right == nil && sum == 0 {
		ans = append(ans, path)
	}
	// 递归左右子树
	dfs(root.Left, sum)
	dfs(root.Right, sum)
	// 不满足条件
	path = path[:len(path)-1]
}