package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
94. Binary Tree Inorder Traversal
todos::用迭代的方式中序遍历二叉树，问题来了，递归和迭代有什么区别
*/
func inorderTraversal(root *TreeNode) []int {
	// 声明空数组
	res := make([]int,0)
	if root == nil{
		return res
	}
	dfs(root,&res)
	return res

}
func dfs(node *TreeNode,res *[]int){
	if node ==nil{
		return
	}
	dfs(node.Left,res)
	// TODOS：：这个什么语法
	(*res) = append(*res, node.Val)
	dfs(node.Right,res)
}