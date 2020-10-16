/*
	根据一棵树的中序遍历与后序遍历构造二叉树。
	tips：
		后序遍历（左、右、根）
		中序遍历（左、根、右）
*/

var pos = make(map[int]int)

func buildTree(inorder []int, postorder []int) *TreeNode {
	for i := 0; i < len(inorder); i++ {
		// 重点：保存中序遍历节点的index
		pos[inorder[i]] = i
	}
	return build(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func build(inorder []int, postorder []int, il, ir, pl, pr int) *TreeNode {
	//
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
