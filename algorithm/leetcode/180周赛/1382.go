/**
	
	1382:将二叉搜索树变平衡

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

	思路：中序遍历（dfs）得到有序数组，再将有序数组转化成平衡二叉树
	todo：证明（难）满足平衡二叉树😅
	todo：提交不通过😅
 */
 var nodes []*TreeNode
 func balanceBST(root *TreeNode) *TreeNode {
	 dfs(root)
	 return build(0,len(nodes)-1)
 }
 func build(l,r int) *TreeNode {
	 if l>r {
		 return nil
	 }
	 mid:=(l+r)/2
	 // todo：关键点
	 nodes[mid].Left = build(l,mid-1)
	 nodes[mid].Right = build(mid+1,r)
	 return nodes[mid]
 }
 // 中序遍历
 func dfs(root *TreeNode)  {
	 if root == nil{
		 return
	 }
	 dfs(root.Left)
	 nodes = append(nodes,root)
	 dfs(root.Right)
 }