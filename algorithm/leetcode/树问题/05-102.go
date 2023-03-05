/*
	思路：借助【队列（先进先出）】实现层序遍历
	给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
*/
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	// 初始化队列
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	// （1）遍历每一层
	for len(queue) > 0 {
		level := make([]int, 0)
		cLen := len(queue)
		// （2）遍历每一层的每个节点
		for i := 0; i < cLen; i++ {
			// 出队
			cnode := queue[0]
			queue = queue[1:]
			// 【根】
			level = append(level, cnode.Val)
			// 【左】
			if cnode.Left != nil {
				queue = append(queue, cnode.Left)
			}
			// 【右】
			if cnode.Right != nil {
				queue = append(queue, cnode.Right)
			}
		}
		res = append(res, level)
	}
	return res
}