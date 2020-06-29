package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	173. Binary Search Tree Iterator
	中序遍历的方式输出这棵树（要求空间和树的高度成正比，用栈模拟中序遍历）
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
	this.stack, top = this.stack[:size-1], this.stack[size-1]
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
		root = root.Left
	}
}
