/*
	给定一棵二叉搜索树，请找出其中的第k小的结点

	知识点：二叉搜索树，中序遍历（左根右）是递增的顺序
	中序遍历：左根右（根节点左边是左子树，右边是右子树）
*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
var ans = &TreeNode{}

func kthNode(root *TreeNode, k int) *TreeNode {
	dfs(root, &k)
	return ans
}

// todo：注意这里要传递指针
func dfs(root *TreeNode, k *int) {
	if root == nil{
		return
	}
	// 遍历左子树
	dfs(root.Left, k)
	// 遍历根节点
	*k--

	if *k == 0{
		ans = root
	}
	
	if *k>0{
		// 遍历右子树
		dfs(root.Right, k)
	}
	
}