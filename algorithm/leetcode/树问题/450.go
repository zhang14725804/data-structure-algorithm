/*
	给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

	一般来说，删除节点可分为两个步骤：
		首先找到需要删除的节点；
		如果找到了，删除它。

	说明： 要求算法时间复杂度为 O(h)，h 为树的高度。
*/

/*
	方法1：DFS-递归
	😅😅😅 val和root.Val三种情况，五种删除逻辑
*/
func deleteNode(root *TreeNode, target int) *TreeNode {
	// （1）没找到删除的节点，遍历到空节点直接返回了
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
		// （3）目标节点同时存在左右节点；必须找到【左子树中最大的节点】，或者【右子树中最小的节点】来接替自己
		// （3.1）找到右子树最小节点
		cnode := root.Right
		for cnode.Left != nil {
			cnode = cnode.Left
		}
		// （3.2）把root改为cnode
		root.Val = cnode.Val
		// （3.3）删除cnode（替换）
		root.Right = deleteNode(root.Right, cnode.Val)
	} else if root.Val > target {
		// 😅 要删除的点在左子树，不是return 哦
		root.Left = deleteNode(root.Left, target)
	} else if root.Val < target {
		// 😅 要删除的点在右子树，不是return 哦
		root.Right = deleteNode(root.Right, target)
	}
	return root
}

/*
	方法2：DFS-迭代
	TODO 😅
*/