/*
	给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
	二叉搜索树：若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值； 若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值
*/

/*
	方法1：递归
*/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(l, r int) []*TreeNode {
	var res []*TreeNode
	// base case
	if l > r {
		res = append(res, nil)
		return res
	}
	for i := l; i <= r; i++ {
		// 递归生成所有左右子树
		left := helper(l, i-1)
		right := helper(i+1, r)
		// 拼接左右子树后返回
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