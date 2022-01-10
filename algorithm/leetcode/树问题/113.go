/*
	给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
*/
/*
	方法1：递归+回溯
	😅😅😅  为什么需要回溯 question
	0110 懵逼状态
*/
var res [][]int
var path []int

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return res
	}
	// 代码提交的时候特殊处理
	res = make([][]int,0)
    path = make([]int,0)
	// (1) 处理根节点
	targetSum -= root.Val
	path = append(path, root.Val)
	backtrack(root, targetSum)
	return res
}

func backtrack(root *TreeNode, targetSum int) {
	// （2）base case 叶子节点，判断是否符合条件“路径上所有节点值相加等于目标和 targetSum”
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		// question 特别注意 😅😅😅😅😅😅😅😅😅 golang切片性质，如果没有这两部，结果是重复的相同数组
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}
	// （3）回溯【左右】子树
	if root.Left != nil {
		// 递归
		targetSum -= root.Left.Val
		path = append(path, root.Left.Val)
		backtrack(root.Left, targetSum)
		// 回溯 😅😅😅
		targetSum += root.Left.Val
		path = path[:len(path)-1]
	}

	if root.Right != nil {
		// 递归
		targetSum -= root.Right.Val
		path = append(path, root.Right.Val)
		backtrack(root.Right, targetSum)
		// 回溯 😅😅😅
		targetSum += root.Right.Val
		path = path[:len(path)-1]
	}
}

