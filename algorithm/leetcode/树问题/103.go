/*
	给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

	思路：层序遍历之后，按照层级反转答案即可
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	step := 0
	// 层序遍历
	for len(q) > 0 {
		ql := len(q)
		var level []int
		for i := 0; i < ql; i++ {
			cur := q[0]
			q = q[1:]
			level = append(level, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		// 根据奇偶反转
		if step == 1 {
			step = 0
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		} else {
			step = 1
		}
		ans = append(ans, level)
	}
	return ans
}