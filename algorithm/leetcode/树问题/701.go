/*
	给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。
	注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
*/

/*
	如何调整搜索二叉树，是个问题。（我想的复杂了 ）😅😅😅😅
	对数据结构的操作无非遍历 + 访问，遍历就是「找」，访问就是「改」。具体到这个问题，插入一个数，就是先找到插入位置，然后进行插入操作。
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	// 递归重建左、右子树
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	// 返回根节点
	return root
}