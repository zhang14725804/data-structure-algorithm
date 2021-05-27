/*
	给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
*/
/*
	方法1：递归+回溯
	😅😅😅 为什么需要回溯 question
*/
var res [][]int
var path []int

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return res
	}
	targetSum -= root.Val
	path = append(path, root.Val)
	dfs(root, targetSum)
	return res
}

func dfs(root *TreeNode, targetSum int) {
	// 叶子节点，判断是否符合条件“路径上所有节点值相加等于目标和 targetSum”
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		// question 特别注意 😅😅😅😅😅😅😅😅😅 golang切片性质，如果没有这两部，结果是重复的相同数组
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}
	if root.Left != nil {
		targetSum -= root.Left.Val
		path = append(path, root.Left.Val)
		// 递归
		dfs(root.Left, targetSum)
		// 回溯 😅😅😅
		targetSum += root.Left.Val
		path = path[:len(path)-1]
	}

	if root.Right != nil {
		targetSum -= root.Right.Val
		path = append(path, root.Right.Val)
		// 递归
		dfs(root.Right, targetSum)
		// 回溯 😅😅😅
		targetSum += root.Right.Val
		path = path[:len(path)-1]
	}
}

