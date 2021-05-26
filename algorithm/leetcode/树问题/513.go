/*
	给定一个二叉树，在树的最后一行找到最左边的值。
*/

/*
	方法1：递归+回溯
	😅，使用递归法，如何判断是最后一行呢，其实就是深度最大的叶子节点一定是最后一行。
*/
var maxLen int = 0   // 记录最大深度
var maxLeftValue int // 记录最大深度左节点的数值
// leftLen 最深深度
func traversal(root *TreeNode, leftLen int) {
	// base case，遇到叶子节点
	if root.Left == nil && root.Right == nil {
		// 更新最大深度的条件
		if leftLen > maxLen {
			maxLen = leftLen        // 更新最大深度
			maxLeftValue = root.Val // 更新最大深度左边的值
		}
		return
	}
	// 😅😅😅 这里隐藏了回溯
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

/*
	方法2：BFS-层序遍历
	最后一行第一个节点的数值
*/
func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	result := 0
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			// 队列先进先出
			node := queue[0]
			queue = queue[1:]
			// 记录最后一行第一个元素
			if i == 0 {
				result = node.Val
			}
			// 先左后右，先进先出
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