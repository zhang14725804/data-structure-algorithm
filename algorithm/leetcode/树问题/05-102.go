/*
	给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
*/

/*
	思路：借助【队列】实现层序遍历
	12.30 面试遇到
*/
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	// 初始化队列
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		level := make([]int, 0)
		// 当前队列长度
		cLen := len(queue)
		for i := 0; i < cLen; i++ {
			// 出队
			cnode := queue[0]
			queue = queue[1:]
			// 【根】
			level = append(level, cnode.Val)
			// 【左右】
			if cnode.Left != nil {
				queue = append(queue, cnode.Left)
			}
			if cnode.Right != nil {
				queue = append(queue, cnode.Right)
			}
		}
		res = append(res, level)
	}
	return res
}