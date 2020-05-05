/*
	输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表
	todo：思路不错
*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Pair struct{
	first *TreeNode
	second *TreeNode
}

func convert(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}
	sides := dfs(root)
	return sides.first
}

func dfs(root *TreeNode) *Pair{
	// 左右子树都为空
	if root.Left ==nil && root.Right ==nil{
		return &Pair{root,root}
	}
	// 左右子树都存在
	if root.Left !=nil && root.Right !=nil{
		lsides := dfs(root.Left)
		rsides := dfs(root.Right)

		lsides.second.Right = root
		root.Left = lsides.second

		root.Right = rsides.first
		rsides.first.Left = root

		return &Pair{lsides.first, rsides.second}
	}
	// 只存在左子树
	if root.Left != nil{
		lsides := dfs(root.Left)

		lsides.second.Right = root
		root.Left = lsides.second

		return &Pair{lsides.first, root}
	}
	// 只存在右子树
	if root.Left != nil{
		rsides := dfs(root.Right)

		root.Right = rsides.first
		rsides.first.Left = root

		return &Pair{root, rsides.second}
	}
	return &Pair{nil,nil}
}
