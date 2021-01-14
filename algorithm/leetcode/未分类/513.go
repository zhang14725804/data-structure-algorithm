/*
	给定一个二叉树，在树的最后一行找到最左边的值。

	方法：递归、迭代
*/

// 方法1：递归+回溯
var maxLen int = 0   // 记录最大深度
var maxLeftValue int // 记录最大深度左节点的数值

func traversal(root *TreeNode, leftLen int) {
	if root.Left == nil && root.Right == nil {
		if leftLen > maxLen {
			maxLen = leftLen
			maxLeftValue = root.Val
		}
		return
	}
	if root.Left != nil {
		traversal(root.Left, leftLen+1)
	}
	if root.Right != nil {
		traversal(root.Right, leftLen+1)
	}
	return
}

func findBottomLeftValue(root *TreeNode) int {
	traversal(root, 0)
	return maxLeftValue
}

// 方法2：迭代，最后一行第一个节点的数值
func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	result := 0
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			// 记录最后一行第一个元素
			if i == 0 {
				result = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}