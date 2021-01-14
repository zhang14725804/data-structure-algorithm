/*
	您需要在二叉树的每一行中找到最大的值。
*/

// 方法1：BFS，遍历每一层，获取最大值
func largestValues(root *TreeNode) []int {
	queue := make([]*TreeNode, 0)
	result := make([]int, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		max := -(1 << 32)
		size := len(queue) // 每一层的数量
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			max = MaxInt(max, node.Val) // 获取较大值
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, max)
	}
	return result
}

// 方法二：dfs；返回结果空数组（question）；没有用指针
func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, res, 1)
	return res
}

// lavel表示层数
func dfs(root *TreeNode, res []int, level int) {
	if root == nil {
		return
	}
	if level == len(res)+1 {
		res = append(res, root.Val)
	} else {
		// level是从1开始的
		res[level-1] = MaxInt(res[level-1], root.Val)
	}
	dfs(root.Left, res, level+1)
	dfs(root.Right, res, level+1)
}

// 方法二：dfs；需要用指针和地址🔥🔥🔥
func dfs(root *TreeNode, res *[]int, level int) {

	if root == nil {
		return
	}

	if level == len(*res)+1 {
		*res = append(*res, root.Val)
	} else {
		// type *[]int not support indexing
		(*res)[level-1] = MaxInt((*res)[level-1], root.Val)
	}
	dfs(root.Left, res, level+1)
	dfs(root.Right, res, level+1)
}
func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, &res, 1)
	return res
}