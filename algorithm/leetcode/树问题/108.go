/*
	将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
	本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

	思路：平衡二叉树等价于中序遍历(左、根、右)是有序的
*/

/*
	方法1：DFS-递归
	前序遍历：根左右
	注意slice取值范围 【左闭右开】（居然卡到这里了） 😅
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 取中间节点
	mid := len(nums) / 2
	// 根
	root := &TreeNode{Val: nums[mid]}
	// 左 【 0 ~ mid-1】
	root.Left = sortedArrayToBST(nums[:mid])
	// 右 【mid+1 ~ len(nums)】
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

/*
	方法2：迭代，三个队列模拟
	question 不理解 😅😅
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 根节点
	root := &TreeNode{}
	// 存放节点
	var nodeQueue []*TreeNode
	// 存放左右区间下标
	var leftQueue, rightQueue []int
	// 根，左右下表入队
	nodeQueue = append(nodeQueue, root)
	leftQueue = append(leftQueue, 0)
	rightQueue = append(rightQueue, len(nums)-1)

	for len(nodeQueue) > 0 {
		cnode := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		left := leftQueue[0]
		leftQueue = leftQueue[1:]
		right := rightQueue[0]
		rightQueue = rightQueue[1:]

		mid := left + (right-left)/2
		cnode.Val = nums[mid]

		// 处理左区间
		if left <= mid-1 {
			cnode.Left = &TreeNode{}
			nodeQueue = append(nodeQueue, cnode.Left)
			leftQueue = append(leftQueue, left)
			rightQueue = append(rightQueue, mid-1)
		}

		// 处理右区间
		if right >= mid+1 {
			cnode.Right = &TreeNode{}
			nodeQueue = append(nodeQueue, cnode.Right)
			leftQueue = append(leftQueue, mid+1)
			rightQueue = append(rightQueue, right)
		}
	}
	return root
}