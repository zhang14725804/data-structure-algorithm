/*
	todo😅：经典题：二叉树前序遍历和中序遍历的结果，重建该二叉树
	前序遍历：根左右（能知道根节点）
	中序遍历：左根右（根节点左边是左子树，右边是右子树）
	后序遍历：左右根
	（还有层序遍历，广度遍历，深度遍历）
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 记录数据引用
var hash = make(map[int]int,0)
var	preorder,inorder []int

func buildTree(_preorder []int, _inorder []int) *TreeNode {
	preorder,inorder = _preorder,_inorder
	// 
	for i:=0;i<len(inorder);i++{
		hash[inorder[i]] = i
	}
	return dfs(0,len(preorder)-1,0,len(inorder)-1)
}

func dfs(pl,pr,il,ir int) *TreeNode{
	if pl > pr {
		return nil
	}
	root := &TreeNode{
		Val:preorder[pl],
		Left:nil,
		Right:nil,
	}
	k:=hash[root.Val]
	// todo:难点，求左右子树
	left:=dfs(pl+1, pl+k-il, il, k-1)
	right:=dfs(pl+k-il+1, pr, k+1, ir)
	root.Left = left
	root.Right = right
	return root
}