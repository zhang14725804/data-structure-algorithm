/*
	给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

	一般来说，删除节点可分为两个步骤：
		首先找到需要删除的节点；
		如果找到了，删除它。

	说明： 要求算法时间复杂度为 O(h)，h 为树的高度。
*/
func deleteNode(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	// 找到，删除。三种情况需要处理(question 重点)
	if root.Val == target {
		// （1）恰好是末端节点，两个子节点都为空，直接删除
		// （2）只有一个非空子节点，那么它要让这个孩子接替自己的位置
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// （3）目标节点同时存在左右节点；必须找到左子树中最大的那个节点，或者右子树中最小的那个节点来接替自己
		minNode := getMinNodeOfBST(root.Right)           // 找到右子树最小节点
		root.Val = minNode.Val                           // 把root改为minNode
		root.Right = deleteNode(root.Right, minNode.Val) // 删除minNode（替换）
	} else if root.Val > target {
		root.Left = deleteNode(root.Left, target)
	} else if root.Val < target {
		root.Right = deleteNode(root.Right, target)
	}
	return root
}

// 找到BST最小节点
func getMinNodeOfBST(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}