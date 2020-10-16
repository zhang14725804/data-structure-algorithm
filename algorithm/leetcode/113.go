/*
	给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

	todo：代码有问题？？
*/
var ans = make([][]int, 0)
var path = make([]int, 0)

func pathSum(root *TreeNode, sum int) [][]int {
	if root != nil {
		dfs(root, sum)
	}
	return ans
}
func dfs(root *TreeNode, sum int) {
	path = append(path, root.Val)
	sum -= root.Val
	// 是叶子节点，而且sum为0
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			ans = append(ans, path)
		}
	} else {
		if root.Left != nil {
			dfs(root.Left, sum)
		}
		if root.Right != nil {
			dfs(root.Right, sum)
		}
	}
	// 恢复现场
	path = path[:len(path)-1]
}