/*
	根据一棵树的中序遍历与后序遍历构造二叉树。
*/

var pos = make(map[int]int)

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	for i := 0; i < n; i++ {
		// 重点：保存中序遍历节点的index
		pos[inorder[i]] = i
	}
	return build(inorder, postorder, 0, n-1, 0, n-1)
}

/*
	inorder   中序遍历数组
	postorder 后序遍历数组
	il，ir    中序遍历起点终点
	pl，pr    后序遍历起点终点
*/
func build(inorder []int, postorder []int, il, ir, pl, pr int) *TreeNode {
	// base case
	if pl > pr {
		return nil
	}
	// 后序遍历的节点值，根节点
	val := postorder[pr]
	// 找到对应中序遍历的index
	k := pos[val]
	// 根节点，对应后序遍历的pr对应的节点
	root := &TreeNode{val, nil, nil}
	// 后序遍历左右子节点的长度和中序遍历左右节点长度相等，ps：k-1-il是左子树长度
	root.Left = build(inorder, postorder, il, k-1, pl, pl+(k-1-il))
	root.Right = build(inorder, postorder, k+1, ir, pl+(k-1-il+1), pr-1)
	return root
}
