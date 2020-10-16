/*
	输入一个二叉树，将它变换为它的镜像。
*/
// 思路：所有左右子树互换位置
func mirror(root *TreeNode) {
	if root == nil {
		return
	}
	mirror(root.Left)
	mirror(root.Right)
	swap(root)
}
func swap(root *TreeNode) {
	root.Left, root.Right = root.Right, root.Left
}