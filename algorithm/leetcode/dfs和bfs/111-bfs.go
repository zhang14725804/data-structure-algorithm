/*
	给定一个二叉树，找出其最小深度。

	最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

	说明: 叶子节点是指没有子节点的节点。
*/

/*
	DFS 实际上是靠递归的堆栈记录走过的路径，你要找到最短路径，肯定得把二叉树中所有树杈都探索完才能对比出最短的路径有多长对不对？而 BFS 借助队列做到一次一步「齐头并进」，是可以在不遍历完整棵树的条件下找到最短距离的。

	形象点说，DFS 是线，BFS 是面；DFS 是单打独斗，BFS 是集体行动。
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 1
	for len(queue) > 0 {
		length := len(queue)
		// 当前队列中的所有节点向四周扩散
		for i := 0; i < length; i++ {
			// ps: 取队头元素
			cur := queue[0]
			queue = queue[1:]
			// 是否到达重点
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			// 左右节点加入队列
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 每一层增加一步
		depth++
	}
	return depth
}
