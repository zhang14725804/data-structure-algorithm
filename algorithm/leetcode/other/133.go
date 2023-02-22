/*
	给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。

	思路：dfs或bfs（需要手动维护队列）
*/
func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}
	table := make(map[int]*Node)
	return dfs(node, table)
}

func dfs(node *Node, table map[int]*Node) *Node {
	if value, ok := table[node.Val]; ok {
		return value
	}

	n := &Node{}
	n.Val = node.Val
	n.Neighbors = make([]*Node, 0)

	table[node.Val] = n
	for _, ni := range node.Neighbors {
		n.Neighbors = append(n.Neighbors, dfs(ni, table))
	}
	return n
}
