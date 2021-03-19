/*
	给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：

	二叉树的根是数组中的最大元素。
	左子树是通过数组中最大值左边部分构造出的最大二叉树。
	右子树是通过数组中最大值右边部分构造出的最大二叉树。
	通过给定的数组构建最大二叉树，并且输出这个树的根节点。
*/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	// 前序遍历构造
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left, right int) *TreeNode {
	// base case
	if left > right {
		return nil
	}
	// 找到最大的元素，和其索引
	index, max := -1, -(1 << 32)
	// 注意临界点
	for i := left; i <= right; i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	root := &TreeNode{Val: max}
	// 递归构造左右子树
	root.Left = build(nums, left, index-1)
	root.Right = build(nums, index+1, right)
	return root
}