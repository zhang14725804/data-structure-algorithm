/*
	二叉树的后序遍历
*/

// 方法1：递归
func postorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)
	return res
}

// 方法1：递归
var ans []int

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return ans
	}
	postorderTraversal(root.Left)
	postorderTraversal(root.Right)
	ans = append(ans, root.Val)
	return ans
}

// 方法2：迭代（TODO 😅）
func postorderTraversal(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	// 通过lastVisit标识右子节点是否已经弹出
	var lastVisit *TreeNode
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 这里先看看，先不弹出
		node := stack[len(stack)-1]
		// 根节点必须在右节点弹出之后，再弹出
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.Val)
			// 标记当前这个节点已经弹出过
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return result
}