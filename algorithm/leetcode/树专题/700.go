/*
	给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	// base case
	if root == nil {
		return nil
	}
	// 三种情况
	if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	}
	// 未找到
	return nil
}