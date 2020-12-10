/*
	给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

	一般来说，删除节点可分为两个步骤：
		首先找到需要删除的节点；
		如果找到了，删除它。

	说明： 要求算法时间复杂度为 O(h)，h 为树的高度。
*/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	// 找到，删除。三种情况需要处理(question 重点)
	if root.Val == key {
		// 目标节点在末尾，或者目标节点之后只有一个元素
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 目标节点同时存在左右节点
		minNode := getMinNode(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

// 找到最小节点
func getMinNode(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}