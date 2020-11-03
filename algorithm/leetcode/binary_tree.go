/*
	二叉树遍历：前序遍历，中序遍历，后序遍历，层序遍历；前中后都是相对于【根元素】而言

	递归都简单，迭代不好理解(question)
	除了递归还有很多实现方法(question)
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans []int

// 递归实现 前序遍历
func preorder(root *TreeNode) []int {
	if root == nil {
		return ans
	}
	dfs := func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

// 迭代实现(需要借助栈，后进先出) 前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return ans
	}
	stk := &Stack{}
	stk.push(root)
	for !stk.empty() {
		cur := stk.pop()
		if cur == nil {
			continue
		}
		ans = append(ans, cur.Val)
		stk.push(cur.Right)
		stk.push(cur.Left)
	}
	return ans
}

// 递归实现 中序遍历
func inorder(root *TreeNode) []int {
	if root == nil {
		return ans
	}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

// 递归实现 后序遍历
func postorderTraversal(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	//
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		ans = append(ans, node.Val)
	}
	dfs(root)
	return ans
}

// 不懂😅，prev什么作用
func postorderTraversal(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode

	for root != nil || len(stk) > 0 {
		// 遍历左子树
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		// 去除栈顶元素
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		// 判断是否存在右子树；root.Right == prev 什么意思（question）
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			// 遍历右子树
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}