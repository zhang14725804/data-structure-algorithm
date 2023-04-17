/*
	给定一个二叉树，找出其最大深度。
	二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
	说明: 叶子节点是指没有子节点的节点。
*/

/*
	方法1：DFS-递归实现
	后序遍历：左右根 😅，question

*/
func maxDepth(root *TreeNode) int {
	// （1）确定递归函数的参数和返回值
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		// （2）base case，终止条件
		if root == nil {
			return 0
		}
		// （3）确定单层递归的逻辑
		ld := dfs(root.Left)  // 左
		rd := dfs(root.Right) //右
		// 😅😅😅 递归下一层，需要 【+1】
		depth := MaxInt(ld, rd) + 1 // 根
		return depth
	}
	return dfs(root)
}

/*
	方法2：BFS-层序遍历
	队列实现 😅😅😅
*/
func maxDepth(root *TreeNode) int {
	//
	if root == nil {
		return 0
	}
	// 初始化queue和level
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		level++
		cLen := len(queue)
		for i := 0; i < cLen; i++ {
			// 出队
			cnode := queue[0]
			queue = queue[1:]
			// 入队
			if cnode.Left != nil {
				queue = append(queue, cnode.Left)
			}
			if cnode.Right != nil {
				queue = append(queue, cnode.Right)
			}
		}
	}
	return level
}

/*
	方法3：BFS-层序遍历
	如何用 栈实现 😅😅😅
	TODO
*/