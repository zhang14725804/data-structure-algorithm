/*
	根据一棵树的前序遍历与中序遍历构造二叉树。

	tips：
		前序遍历（根、左、右）
		中序遍历（左、根、右）
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// map 保存中序遍历节点的index
var pos = make(map[int]int)

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(inorder)
	for i := 0; i < n; i++ {
		// 重点：保存中序遍历节点的index
		pos[inorder[i]] = i
	}
	return dfs(preorder, inorder, 0, n-1, 0, n-1)
}

/*
	preorder，前序遍历数组
	inorder，中序遍历数组
	pl，pr，il，ir
*/
func dfs(preorder, inorder []int, pl, pr, il, ir int) *TreeNode {
	//
	if pl > pr {
		return nil
	}
	// 前序遍历的节点值，根节点
	val := preorder[pl]
	// 找到对应中序遍历的index
	k := pos[val]
	// 根节点，对应前序遍历的pl对应的节点
	root := &TreeNode{val, nil, nil}
	/*
		pl+1, pl+(k-il), il, k-1。ps:k-il是左子树长度
	*/
	root.Left = dfs(preorder, inorder, pl+1, pl+k-il, il, k-1)
	/*
		pl+(k-il)+1, pr, k+1, ir
	*/
	root.Right = dfs(preorder, inorder, pl+k-il+1, pr, k+1, ir)
	return root
}
