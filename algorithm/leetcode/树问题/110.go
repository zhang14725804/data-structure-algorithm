/*
	给定一个二叉树，判断它是否是高度平衡的二叉树。

	本题中，一棵高度平衡二叉树定义为：

	一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
*/

/*
	方法1：DFS-递归，自顶向下方式，可以缓存节点高度 😅
	问题：因此对于同一个节点，函数 height 会被重复调用，导致时间复杂度较高。
	如果使用自底向上的做法，则对于每个节点，函数 height 只会被调用一次
*/
func isBalanced(root *TreeNode) bool {
	// base case
	if root == nil {
		return true
	}
	// 😅 返回左右节点最大高度
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	// 判断左右节点是否是高度平衡的二叉树
	if abs(left, right) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

/*
	方法2：DFS-递归，自底向上 😅

	question : 思路不错
	分治法，左边平衡 && 右边平衡 && 左右两边高度 <= 1，
	因为需要返回是否平衡及高度，要么返回两个数据，要么合并两个数据，
	所以用 【-1】 表示不平衡，【>0】 表示树高度（二义性：一个变量有两种含义 😅😅😅）
*/
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

// 后序遍历：左右根
func height(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}
	left := height(root.Left)   // 左
	right := height(root.Right) // 右
	if left == -1 || right == -1 || abs(left, right) > 1 {
		return -1
	}
	// 根，返回当前二叉树高度
	return MaxInt(left, right) + 1
}