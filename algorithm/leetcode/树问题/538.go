/*
	给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
*/

/*
	方法1：DFS-递归
	从树中可以看出累加的顺序是【右根左】😅😅😅，所以我们需要反中序遍历这个二叉树，然后顺序累加就可以了。
*/

var sum int // 记录前一个节点的数值

func convertBST(root *TreeNode) *TreeNode {
	dfs(root)
	return root
}

// 反中序遍历：右根左 😅😅😅
func dfs(root *TreeNode) {
	// base case
	if root == nil {
		return
	}
	// 确定单层递归的逻辑
	dfs(root.Right)
	// 😅 先计算sum，再赋值
	sum += root.Val
	root.Val = sum
	dfs(root.Left)
}

/*
	方法2：DFS-迭代
*/
var sum int

func convertBST(root *TreeNode) *TreeNode {
	helper(root)
	return root
}

func helper(root *TreeNode) {
	var stack []*TreeNode
	cnode := root
	for cnode != nil || len(stack) > 0 {
		if cnode != nil {
			stack = append(stack, cnode)
			cnode = cnode.Right
		} else {
			cLen := len(stack) - 1
			cnode = stack[cLen]
			stack = stack[:cLen]
			// 😅 先计算sum，再赋值
			sum += cnode.Val
			cnode.Val = sum

			cnode = cnode.Left
		}
	}
}