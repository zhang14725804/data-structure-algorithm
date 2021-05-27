/*
	给你二叉树的根节点 root 和一个表示目标和的整数 targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
	叶子节点 是指没有子节点的节点。
*/

/*
	方法1：DFS-递归+回溯
	question 隐藏了回溯的细节 😅😅😅
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	// 😅😅 base case，判断是否满足条件，我忽略了最重要的这一步
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	// 递归判断下一层
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

/*
	方法2：DFS-迭代，栈模拟递归 😅😅😅
*/
// 栈里一个元素不仅要记录该节点指针，还要记录从头结点到该节点的路径数值总和
type Pair struct {
	First  *TreeNode
	Second int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var stack []*Pair
	stack = append(stack, &Pair{First: root, Second: root.Val})
	for len(stack) > 0 {
		cLen := len(stack) - 1

		cnode := stack[cLen]
		stack = stack[:cLen]

		/*
			这么写只判断一次叶子节点 😅😅😅
			if cnode.First.Left == nil && cnode.First.Right == nil {
				return cnode.Second == targetSum
			}
		*/
		// 叶子节点，判断是否符合条件“路径上所有节点值相加等于目标和 targetSum”
		if cnode.First.Left == nil && cnode.First.Right == nil && cnode.Second == targetSum {
			return true
		}
		if cnode.First.Right != nil {
			stack = append(stack, &Pair{First: cnode.First.Right, Second: cnode.Second + cnode.First.Right.Val})
		}
		if cnode.First.Left != nil {
			stack = append(stack, &Pair{First: cnode.First.Left, Second: cnode.Second + cnode.First.Left.Val})
		}

	}
	return false
}