/*
	实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
	调用 next() 将返回二叉搜索树中的下一个最小的数。
	中序遍历的方式输出这棵树（要求空间和树的高度成正比，用栈模拟中序遍历）

	next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
	你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。

	todo：这个题蒙蔽（用栈模拟中序遍历：左根右）
*/

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	// todos：这里为什么这么写
	stack := make([]*TreeNode, 0, 128)
	res := BSTIterator{
		stack: stack,
	}
	res.push(root)
	return res
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	size := len(this.stack)
	var top *TreeNode
	// top、pop
	this.stack, top = this.stack[:size-1], this.stack[size-1]
	// 右子树加入栈
	this.push(top.Right)
	return top.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

func (this *BSTIterator) push(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		// 左子树加入栈
		root = root.Left
	}
}
