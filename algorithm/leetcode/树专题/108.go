/*
	将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
	本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

	思路：平衡二叉树等价于中序遍历(左、根、右)是有序的

	测试不通过，提交通过了,,ԾㅂԾ,,
*/
func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

//
func build(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	// 根节点
	mid := (left + right) >> 1
	root := &TreeNode{Val: nums[mid]}
	// 左右子节点
	root.Left = build(nums, left, mid-1)
	root.Right = build(nums, mid+1, right)
	return root
}