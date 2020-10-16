/*
	二叉树前序遍历（根左右）

	容器属性：flex-wrap,flex-direction,flex-flow,align-items,justify-content,align-content
	元素属性：order,flex-grow,flex-shrink,flex-basis,align-self,flex
*/

var ans []int

// 方法一：递归
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return ans
	}
	dfs(root)
	return ans
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	ans = append(ans, root.Val)
	dfs(root.Left)
	dfs(root.Right)
}

// 方法二：迭代
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return ans
	}
	stack := &Stack{}
	stack.push(root)
	for !stack.empty() {
		cur := stack.pop()
		if cur == nil {
			continue
		}
		ans = append(ans, cur.Val)
		stack.push(cur.Right)
		stack.push(cur.Left)
	}
	return ans
}

