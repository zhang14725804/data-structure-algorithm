/*
	给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
	依次输出二叉树每层最右边的元素。

	方法1： 根右左的顺序递归，每次取第一次到达改层的元素（ len(res) 表示当前几层 ）
	方法2： 根左右的顺序层序遍历，每次取每一层最后一个元素
	0105 懵逼
*/

/*
	方法1：【根右左】 😅😅😅 的顺序递归
	【根右左】的顺序，用一个变量记录当前层数，每次保存第一次到达该层的元素。
*/
func rightSideView(root *TreeNode) []int {
	var res []int
	var dfs func(root *TreeNode, level int)
	// queue 用一个变量记录当前层数，每次保存第一次到达该层的元素。
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 😅 保存第一次到达该层的元素
		if level == len(res) {
			res = append(res, root.Val)
		}
		// 😅 根右左的顺序
		dfs(root.Right, level+1)
		dfs(root.Left, level+1)
	}
	dfs(root, 0)
	return res
}

/*
	思路1：层序遍历【跟左右】的顺序，每次取最后一个元素
*/
func rightSideView(root *TreeNode) []int {
	var res []int
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		cLen := len(queue)
		for i := 0; i < cLen; i++ {
			// 😅 出队动作不能停
			cnode := queue[0]
			queue = queue[1:]
			// 😅 将每一层的最后元素放入result数组中
			if i == cLen-1 {
				res = append(res, cnode.Val)
			}
			//
			if cnode.Left != nil {
				queue = append(queue, cnode.Left)
			}
			if cnode.Right != nil {
				queue = append(queue, cnode.Right)
			}
		}
	}
	return res
}