/*
	给出一个完全二叉树，求出该树的节点个数。

	知识点：完美二叉树、【完全二叉树】、【满二叉树】
*/

/*
	方法1：递归
*/
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

/*
	方法2：利用完全二叉树性质
	对于 complete binary tree ，左子树和右子树中一定存在 perfect binary tree。
	如果当前二叉树是一个 perfect binary tree，我们完全可以用公式算出当前二叉树的总节点数。
	算法复杂度 O(logN*logN)
*/
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 从左右出发，记录左右子树高度
	left := root
	lh := 0
	for left != nil {
		lh++
		left = left.Left
	}
	right := root
	rh := 0
	for right != nil {
		rh++
		right = right.Right
	}
	// 左右子树高度相同，满二叉树
	if lh == rh {
		return (1 << lh) - 1
	} else {
		return countNodes(root.Left) + countNodes(root.Right) + 1
	}
}

/*
	方法3：（没理解）
	如果右子树的高度等于整个树的高度减 1，说明左边都填满了，所以左子树是 perfect binary tree
	否则的话，右子树是 perfect binary tree
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	h := getHeight(root)
	rh := getHeight(root.Right)

	if h-1 == rh {
		return (1 << rh) - 1 + countNodes(root.Right) + 1
	} else {
		return (1 << rh) - 1 + countNodes(root.Left) + 1
	}
}

// 求完全二叉树高度
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getHeight(root.Left) + 1
}