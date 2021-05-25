/*
	给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。
*/
/*
	思路：借助【队列】实现层序遍历，计算每一层的平均值存入res结果集
*/
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		var levelSum float64
		cLen := len(queue)
		for i := 0; i < cLen; i++ {
			// 队列，先进先出
			cnode := queue[0]
			queue = queue[1:]
			levelSum += float64(cnode.Val)
			if cnode.Left != nil {
				queue = append(queue, cnode.Left)
			}
			if cnode.Right != nil {
				queue = append(queue, cnode.Right)
			}
		}
		res = append(res, levelSum/float64(cLen))
	}
	return res
}