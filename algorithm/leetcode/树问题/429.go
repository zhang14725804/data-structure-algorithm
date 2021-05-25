/*
	给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
	树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
*/

/*
	思路：借助【队列】实现层序遍历
*/
func levelOrder(root *Node) [][]int {
	var res [][]int
	var queue []*Node
	if root != nil {
		queue = append(queue, root)
	}
	// 😅😅😅 用队列（先进先出）辅助数据结构
	for len(queue) > 0 {
		level := make([]int, 0)
		cLen := len(queue)
		for i := 0; i < cLen; i++ {
			cnode := queue[0]
			queue = queue[1:]
			// 不需要考虑分割nil值嘛 😅
			level = append(level, cnode.Val)
			if cnode.Children != nil {
				queue = append(queue, cnode.Children...)
			}
		}
		res = append(res, level)
	}
	return res
}