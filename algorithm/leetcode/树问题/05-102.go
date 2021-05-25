/*
	给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
*/

/*
	思路：借助【队列】实现层序遍历 🔥🔥🔥
*/
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		level := make([]int, 0)
		cLen := len(queue)
		for i := 0; i < cLen; i++ {
			// 😅😅😅 用队列（先进先出）辅助数据结构
			cnode := queue[0]
			queue = queue[1:]
			level = append(level, cnode.Val)
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