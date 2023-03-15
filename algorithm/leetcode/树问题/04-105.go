var pos = make(map[int]int, 0)

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 保存【中序遍历】节点的index
	for idx, v := range inorder {
		pos[v] = idx
	}

	n := len(preorder)
	// 本质上是前序遍历：根左右(🔥🔥🔥)
	return dfs(preorder, inorder, 0, n-1, 0, n-1)
}

/*
	preorder，前序遍历数组
	inorder，中序遍历数组
	pl，pr 前序遍历起点、终点
	il，ir 中序遍历起点、终点
*/
func dfs(preorder, inorder []int, pl, pr, il, ir int) *TreeNode {
	// base case
	if pl > pr {
		return nil
	}

	// 跟节点
	val := preorder[pl]
	node := &TreeNode{val, nil, nil}
	// 中序遍历根节点位置
	idx := pos[val]

	/*
		左子树：取前序遍历左边，取中序遍历左边。ps:【idx-il】 是左子树长度(🔥🔥🔥)
		pl+1,
		pl+(idx-il),
		il,
		idx-1。
	*/
	node.Left = dfs(preorder, inorder, pl+1, pl+(idx-il), il, idx-1)
	/*
		右子树：取前序遍历右边，取中序遍历右边。ps:【idx-il】 是左子树长度(🔥🔥🔥)
		pl+(idx-il)+1,
		pr,
		idx+1,
		ir
	*/
	node.Right = dfs(preorder, inorder, pl+(idx-il)+1, pr, idx+1, ir)
	return node
}