/*
	二叉树前序遍历（根左右）
*/

var ans []int

// 方法1：递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
	return ans
}

// 迭代实现前序遍历
var ans []int

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		ans = append(ans, node.Val)
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return ans
}

