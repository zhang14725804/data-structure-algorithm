/*
	给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。
	注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
*/

/*
	方法1：DFS-递归
	如何调整搜索二叉树，是个问题。
	（我想的复杂了 ）😅😅😅😅只要遍历二叉搜索树，找到空节点 插入元素就可以了，那么这道题其实就简单了。
	对数据结构的操作无非遍历 + 访问，遍历就是「找」，访问就是「改」。具体到这个问题，插入一个数，就是先找到插入位置，然后进行插入操作。
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 😅 base case
	if root == nil {
		return &TreeNode{Val: val}
	}
	// 😅 根据BST性质，递归重建左、右子树
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	// 😅 返回根节点
	return root
}

/*
	方法2：迭代 😅😅😅
	在迭代法遍历的过程中，需要记录一下当前遍历的节点的父节点，这样才能做插入节点的操作。
*/
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	cnode := root
	// 😅😅😅 这个很重要，需要记录上一个节点，否则无法赋值新节点
	parent := root
	for cnode != nil {
		parent = cnode
		if cnode.Val > val {
			cnode = cnode.Left
		} else if cnode.Val < val {
			cnode = cnode.Right
		}
	}
	node := &TreeNode{Val: val}
	// 😅😅😅 此时是用parent节点的进行赋值
	if val < parent.Val {
		parent.Left = node
	} else {
		parent.Right = node
	}
	return root
}