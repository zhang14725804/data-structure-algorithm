/*
	给定一个 N 叉树，找到其最大深度。

	最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。

	N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。
*/
/*
	方法1：BFS-层序遍历
*/
func maxDepth(root *Node) int {
	var queue []*Node
	if root != nil {
		queue = append(queue, root)
	}
	var level int
	for len(queue) > 0 {
		// 遍历每一层
		cLen := len(queue)
		for i := 0; i < cLen; i++ {
			cnode := queue[0]
			queue = queue[1:]
			if cnode.Children != nil {
				queue = append(queue, cnode.Children...)
			}
		}
		// 最大深度+1
		level++
	}
	return level
}

/*
	方法2：DFS-递归实现
*/
func maxDepth(root *Node) int {
	// (1) 确定递归函数的参数和返回值
	var dfs func(root *Node) int
	dfs = func(root *Node) int {
		// (2) 确定终止条件 base case
		if root == nil {
			return 0
		}
		// (3) 确定单层递归的逻辑
		var depth int
		for i := 0; i < len(root.Children); i++ {
			depth = MaxInt(depth, dfs(root.Children[i]))
		}
		return depth + 1
	}
	return dfs(root)
}