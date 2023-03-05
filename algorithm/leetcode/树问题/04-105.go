/*
	根据一棵树的前序遍历与中序遍历构造二叉树。
*/

// map 保存【中序遍历】节点的index 😅😅😅
var pos = make(map[int]int)

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	// 😅 保存中序遍历节点的index
	for i := 0; i < n; i++ {
		pos[inorder[i]] = i
	}
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
	// base case 递归出口 😅
	if pl > pr {
		return nil
	}
	// 前序遍历最左边的节点，根节点
	val := preorder[pl]
	// 找到对应中序遍历的index
	k := pos[val]
	// 根节点，对应前序遍历的pl对应的节点
	root := &TreeNode{val, nil, nil}
	/*
		构造【左子树】，中序遍历根据【k】值分割，前序遍历较难分割
		ps:【k-il】 是左子树长度(🔥🔥🔥)
		pl+1,
		pl+(k-il),
		il,
		k-1。
	*/
	root.Left = dfs(preorder, inorder, pl+1, pl+(k-il), il, k-1)
	/*
		构造【右子树】
		ps:【k-il】 是左子树长度(🔥🔥🔥)
		pl+(k-il)+1,
		pr,
		k+1,
		ir
	*/
	root.Right = dfs(preorder, inorder, pl+(k-il)+1, pr, k+1, ir)
	return root
}
