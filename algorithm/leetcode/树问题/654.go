/*
	给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：

	二叉树的根是数组中的最大元素。
	左子树是通过数组中最大值左边部分构造出的最大二叉树。
	右子树是通过数组中最大值右边部分构造出的最大二叉树。
	通过给定的数组构建最大二叉树，并且输出这个树的根节点。
*/

/*
	方法1：DFS-递归
*/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	nLen := len(nums)
	// base case
	if nLen == 0 {
		return nil
	}
	// 找到最大值和最大值索引
	midx := 0
	max := nums[0]
	for i := 1; i < nLen; i++ {
		if max < nums[i] {
			max = nums[i]
			midx = i
		}
	}
	// 😁 前序遍历：根左右
	root := &TreeNode{max, nil, nil}
	// 根据最大值索引切割左右子树
	root.Left = constructMaximumBinaryTree(nums[:midx])
	root.Right = constructMaximumBinaryTree(nums[midx+1:])
	return root
}