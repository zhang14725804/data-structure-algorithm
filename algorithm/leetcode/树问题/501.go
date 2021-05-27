/*
	给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
	你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
*/

/*
	方法1：遍历二叉树+map统计频率（用了额外的空间 😅）
*/

/*
	方法2：DFS-递归
	利用二叉搜索树性质，【中序遍历】是排序数组
	question 😅😅😅
*/
var maxCount, count int
var res []int
var prev *TreeNode

func findMode(root *TreeNode) []int {
	dfs(root)
	return res
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	// 左
	dfs(root.Left)
	// 根
	// 😅 更新count
	if prev == nil { //第一个节点
		count = 1
	} else if prev.Val == root.Val { // 与前一个节点数值相同
		count++
	} else { // 与前一个节点数值不同
		count = 1
	}

	// 😅 更新maxCount
	// 如果和最大值相同，放进result中
	if count == maxCount {
		res = append(res, root.Val)
	}
	// 如果计数大于最大值频率
	if count > maxCount {
		// 更新最大频率
		maxCount = count
		// 😅 清空当前结果数组
		res = make([]int, 0)
		res = append(res, root.Val)
	}
	prev = root
	// 右
	dfs(root.Right)
}

/*
	方法2：DFS-迭代法
	利用二叉搜索树性质，【中序遍历】是排序数组
	question： 没懂 😅😅😅
*/
func findMode(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	cnode := root
	var prev *TreeNode
	var maxCount, count int
	// 为什么不直接把root 放入stack
	for cnode != nil || len(stack) > 0 {
		if cnode != nil {
			stack = append(stack, cnode)
			cnode = cnode.Left
		} else {
			cLen := len(stack) - 1
			cnode = stack[cLen]
			stack = stack[:cLen]

			if prev == nil {
				count = 1
			} else if prev.Val == cnode.Val {
				count++
			} else {
				count = 1
			}
			if count == maxCount {
				res = append(res, cnode.Val)
			}
			if count > maxCount {
				maxCount = count
				res = make([]int, 0)
				res = append(res, cnode.Val)
			}

			prev = cnode
			cnode = cnode.Right
		}
	}
	return res
}