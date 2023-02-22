/*
	您需要在二叉树的每一行中找到最大的值。
*/

/*
	思路1：借助【队列】实现层序遍历 🔥🔥🔥
*/
func largestValues(root *TreeNode) []int {
	var res []int
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		cLen := len(queue)
		// 😅负数的情况
		max := -(1 << 32)
		for i := 0; i < cLen; i++ {
			cnode := queue[0]
			queue = queue[1:]
			max = MaxInt(max, cnode.Val)
			if cnode.Left != nil {
				queue = append(queue, cnode.Left)
			}
			if cnode.Right != nil {
				queue = append(queue, cnode.Right)
			}
		}
		res = append(res, max)
	}
	return res
}

/*
	方法2：递归
*/
func largestValues(root *TreeNode) []int {
	var res []int
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 第一个元素直接加入结果集
		if level == len(res) {
			res = append(res, root.Val)
		} else {
			// 后续元素取较大值
			res[level] = MaxInt(res[level], root.Val)
		}
		// 递归下一层
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return res
}