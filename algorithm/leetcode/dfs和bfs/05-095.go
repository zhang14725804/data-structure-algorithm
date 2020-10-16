/*
	给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
	二叉搜索树：若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值； 若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值
*/
func generateTrees(n int) []*TreeNode {
	//
	if n == 0 {
		return nil
	}
	return dfs(1, n)
}

func dfs(l, r int) []*TreeNode {
	//
	var res []*TreeNode
	if l > r {
		res = append(res, nil)
		return res
	}
	for i := l; i <= r; i++ {
		// 左右子树
		left := dfs(l, i-1)
		right := dfs(i+1, r)
		// 组合左右子树
		for _, lt := range left {
			for _, rt := range right {
				// 生成新节点
				root := &TreeNode{Val: i}
				root.Left = lt
				root.Right = rt
				res = append(res, root)
			}
		}
	}
	return res
}