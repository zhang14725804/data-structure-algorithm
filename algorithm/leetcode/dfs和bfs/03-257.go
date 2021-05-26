/*
	方法1：DFS-递归
	题目要求从根节点到叶子的路径，所以需要【前序遍历】，这样才方便让父节点指向孩子节点，找到对应的路径。
	question 还是没做出来 😅😅😅
	https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0257.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%89%80%E6%9C%89%E8%B7%AF%E5%BE%84.md
*/
var ans []string

func binaryTreePaths(root *TreeNode) []string {
	var path string
	dfs(root, path)
	return ans
}

// 😅😅😅 递归中隐藏了【回溯】
func dfs(root *TreeNode, path string) {
	// base case
	if root == nil {
		return
	}
	// 排除第一个节点的情况，而不是 【path += "->" + fmt.Sprintf("%v", root.Val)】这种写法
	if len(path) > 0 {
		path += "->"
	}
	path += fmt.Sprintf("%v", root.Val) // 根
	if root.Left == nil && root.Right == nil {
		// 叶子节点，添加当前路径
		ans = append(ans, path)
	} else {
		dfs(root.Left, path)  // 左
		dfs(root.Right, path) // 右
	}
}

/*
	方法2：DFS-迭代方式
	稍微有点懵逼 😅😅
*/
func binaryTreePaths(root *TreeNode) []string {
	// 【栈】模拟递归
	var nodeStack []*TreeNode
	// 【栈】来存放对应的遍历路径
	var pathStack []string
	var res []string

	if root == nil {
		return res
	}
	nodeStack = append(nodeStack, root)
	pathStack = append(pathStack, fmt.Sprintf("%v", root.Val))
	for len(nodeStack) > 0 {
		// 栈，先进后出
		nLen := len(nodeStack) - 1
		cnode := nodeStack[nLen]
		nodeStack = nodeStack[:nLen]

		pLen := len(pathStack) - 1
		cpath := pathStack[pLen]
		pathStack = pathStack[:pLen]
		// 叶子节点，添加当前路径
		if cnode.Left == nil && cnode.Right == nil {
			res = append(res, cpath)
		}
		// 栈，先进后出
		if cnode.Right != nil {
			nodeStack = append(nodeStack, cnode.Right)
			pathStack = append(pathStack, cpath+"->"+fmt.Sprintf("%v", cnode.Right.Val))
		}
		if cnode.Left != nil {
			nodeStack = append(nodeStack, cnode.Left)
			pathStack = append(pathStack, cpath+"->"+fmt.Sprintf("%v", cnode.Left.Val))
		}
	}
	return res
}