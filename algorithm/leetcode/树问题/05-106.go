var pos = make(map[int]int, 0)

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 保存【中序遍历】节点的index
	for idx, v := range inorder {
		pos[v] = idx
	}

	n := len(inorder)
	// 本质上是前序遍历：根左右(🔥🔥🔥)
	return dfs(inorder, postorder, 0, n-1, 0, n-1)
}

/*
	inorder   中序遍历数组
	postorder 后序遍历数组
	il，ir    中序遍历起点终点
	pl，pr    后序遍历起点终点
*/
func dfs(inorder, postorder []int, il, ir, pl, pr int) *TreeNode {
	// base case
	if il > ir {
		return nil
	}

	// 根节点
	val := postorder[pr]
	node := &TreeNode{val, nil, nil}
	// 中序遍历根节点位置
	idx := pos[val]

	/*
		左子树：取后序遍历左边，取中序遍历左边。ps:【ir-idx】 是右子树长度(🔥🔥🔥)
		il,
		idx-1,
		pl,
		pr-(ir-idx)-1
	*/
	node.Left = dfs(inorder, postorder, il, idx-1, pl, pr-(ir-idx)-1)

	/*
		右子树：取后序遍历右边，取中序遍历右边。ps:【ir-idx】 是右子树长度(🔥🔥🔥)
		idx+1,
		ir,
		pr-(ir-idx),
		pr-1
	*/
	node.Right = dfs(inorder, postorder, idx+1, ir, pr-(ir-idx), pr-1)

	return node
}