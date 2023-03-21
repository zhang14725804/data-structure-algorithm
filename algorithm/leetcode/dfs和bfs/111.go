/*
	方法1：DFS-递归
	后序遍历：左右根
	😅 求二叉树的【最小深度】和求二叉树的【最大深度】的差别主要在于处理左右孩子不为空的逻辑
	question，再次遇到还是不会，😅 😅 😅 😅
*/
func minDepth(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}
	// 计算左右子树
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	// 😅😅 左边或者右边是空，【+1】
	if left == 0 || right == 0 {
		return left + right + 1
	}
	// 😅😅 左右子树都不为空，【+1】
	return MinInt(left, right) + 1
}

/*
	方法2：BFS-迭代法
	😅😅😅
	DFS 实际上是靠递归的堆栈记录走过的路径，你要找到最短路径，肯定得把二叉树中所有树杈都探索完才能对比出最短的路径有多长对不对？
	而 BFS 借助队列做到一次一步「齐头并进」，是可以在不遍历完整棵树的条件下找到最短距离的。
	形象点说，DFS 是线，BFS 是面；DFS 是单打独斗，BFS 是集体行动。
*/
func minDepth(root *TreeNode) int {
	var queue []*TreeNode
	var depth int
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		// 每一层深度【+1】
		depth++
		cLen := len(queue)
		for i := 0; i < cLen; i++ {
			cnode := queue[0]
			queue = queue[1:]
			// 左右节点加入队列
			if cnode.Left != nil {
				queue = append(queue, cnode.Left)
			}
			if cnode.Right != nil {
				queue = append(queue, cnode.Right)
			}
			// 😅😅😅 当左右孩子都为空的时候，说明是最低点的一层了，退出
			if cnode.Left == nil && cnode.Right == nil {
				return depth
			}
		}
	}
	return depth
}