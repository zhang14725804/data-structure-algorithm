/*
	实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
	调用 next() 将返回二叉搜索树中的下一个最小的数。

	二叉搜索树：左子树 < 根节点 < 右子树
*/

type BSTIterator struct {
	nums []int
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	nums := make([]int, 0)
	inorder(root, &nums) // 传递地址

	return BSTIterator{
		nums: nums,
		root: root,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	res := this.nums[0]
	this.nums = this.nums[1:]
	return res
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.nums) > 0
}

// 中序遍历：左根右
func inorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, nums)
	*nums = append(*nums, root.Val) // 操作指针
	inorder(root.Right, nums)
}