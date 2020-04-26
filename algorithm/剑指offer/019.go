/*
	给定一棵二叉树的其中一个节点，请找出中序遍历序列的下一个节点。

	todo：后继：当前节点在中序遍历中的下一个节点，也就是整颗二叉搜索树中比当前节点大的最小的元素😅
	todo：两种情况
*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
    Father *TreeNode
}

func inorderSuccessor(p *TreeNode) *TreeNode {
	// 有右子树
	if p.Right != nil{
		p = p.Right
		for p.Left != nil{
			p = p.Left
		}
		return p
	}
	// 没有右子树
	for p.Father != nil && p == p.Father.Right{
		p = p.Father
	}
	return p.Father
}