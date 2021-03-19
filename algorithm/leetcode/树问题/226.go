/*
	翻转一棵二叉树。
*/

/*
	方法1：递归（自顶向下）
*/
func invertTree1(root *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return root
	}
	// 前序遍历位置
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	// 递归左右子树
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

/*
	方法2：迭代
	(question)
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		// 从队首获取元素，同时获取的这个元素将从原队列删除；
		temp := queue[0]
		queue = queue[1:]
		// 交换左右子树
		left := temp.Left
		temp.Left = temp.Right
		temp.Right = left

		if temp.Left != nil {
			queue = append(queue, temp.Left)
		}
		if temp.Right != nil {
			queue = append(queue, temp.Right)
		}
	}
	return root
}
