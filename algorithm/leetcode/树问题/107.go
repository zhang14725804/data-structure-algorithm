/*
	给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
*/
/*
	思路：借助【队列】实现层序遍历，然后将层序遍历的结果反转即可
*/
func levelOrderBottom(root *TreeNode) [][]int {
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
	// 反转结果
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}