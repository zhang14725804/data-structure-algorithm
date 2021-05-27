/*
	给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
*/

/*
	方法1：DFS-递归
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	// base case
	if root == nil || root.Val == val {
		return root
	}
	// 前序遍历：根左右，这里居然懵住了 😅
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

/*
	方法2：迭代法 😅😅😅
	根本不需要【使用栈来模拟深度遍历，使用队列来模拟广度遍历】
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}