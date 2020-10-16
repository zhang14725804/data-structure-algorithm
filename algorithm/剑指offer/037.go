/*
	输入两棵二叉树A，B，判断B是不是A的子结构（类似子字符串问题）
*/
func hasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	if isPart(pRoot1, pRoot2) {
		return true
	}
	return hasSubtree(pRoot1.Left, pRoot2) || hasSubtree(pRoot1.Right, pRoot2)
}

func isPart(p1 *TreeNode, p2 *TreeNode) bool {
	// 第二棵树已经遍历完
	if p2 == nil {
		return true
	}
	// 不匹配
	if p1 == nil || p1.Val != p2.Val {
		return false
	}
	// 左右子树都要满足
	return isPart(p1.Left, p2.Left) && isPart(p1.Right, p2.Right)
}