/*
	给你一棵所有节点为非负值的 **二叉搜索树**，请你计算树中任意两节点的差的绝对值的最小值。
	二叉搜索树性值：二叉搜索树中序遍历得到的值序列是递增有序的
	方法：中序遍历即可
*/

/*
	方法1：前序遍历
	前序遍历，在一个有序数组上求两个数最小差值
	🔥🔥🔥 遇到在二叉搜索树上求什么最值啊，差值之类的，就把它想成在一个有序数组上求最值，求差值，这样就简单多了
*/

/*
	方法2：DFS-递归
	利用搜索二叉树性质：【中序遍历】是有序的
*/
var res = (1 << 32) // 声明极大值
var prev *TreeNode

func getMinimumDifference(root *TreeNode) int {
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	// base case
	if root == nil {
		return
	}
	// 😅 前序遍历
	dfs(root.Left)   // 左
	if prev != nil { // 根
		res = MinInt(res, root.Val-prev.Val)
	}
	prev = root
	dfs(root.Right) // 右
}

/*
	方法3：递归-迭代实现
	利用搜索二叉树性质：【中序遍历】是有序的
	question： 没懂 😅😅😅
*/
func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	if root == nil {
		return 0
	}
	var res = (1 << 32)
	cnode := root
	var stack []*TreeNode
	// 为什么不直接把root 放入stack
	for cnode != nil || len(stack) > 0 {
		if cnode != nil {
			stack = append(stack, cnode)
			// 左
			cnode = cnode.Left
		} else {
			cLen := len(stack) - 1
			cnode = stack[cLen]
			stack = stack[:cLen]
			// 根
			if prev != nil {
				res = MinInt(res, cnode.Val-prev.Val)
			}
			// 跟新prev
			prev = cnode
			// 右
			cnode = cnode.Right
		}
	}
	return res
}