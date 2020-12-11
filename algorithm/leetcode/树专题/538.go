/*
	给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
*/

/*
	方法1：逆序遍历BST(question 机智啊)
		BST 的中序遍历代码可以升序打印节点的值
		如果想降序打印节点的值怎么办：把遍历顺序改为  右根左
*/
func convertBST(root *TreeNode) *TreeNode {
	traverse(root)
	return root
}

var sum int

func traverse(root *TreeNode) {
	// base case
	if root == nil {
		return
	}
	traverse(root.Right)
	sum += root.Val
	root.Val = sum
	traverse(root.Left)
}