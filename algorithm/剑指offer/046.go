/*
	输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果
	todo：有难度（和18题重建二叉树类似）
	知识点：
	（1）二叉搜索树（左子树 < 中间节点 < 右子树），
	（2）树的后序遍历（左右根），最后一个节点是根节点

	后续遍历结果：[4, 8, 6, 12, 16, 14, 10]
	对应二叉搜索树：

		10
	6		14	
  4	  8   12   16
*/
var seq []int
func verifySequenceOfBST(sequence []int) bool{
	seq = sequence
	// 递归遍历
	return dfs(0,len(seq)-1)
}

func dfs(l,r int) bool{
	// 已经遍历完区间
	if l>=r{
		return true
	}
	// 当前区间根节点
	root := seq[r]
	k := l
	// 左半边
	for k < r && seq[k] < root{
		k++
	}
	for i := k;i<r;i++{
		// 右半边应该都大于当前根节点
		if seq[i]<root{
			return false
		}
	}
	// 递归遍历左右区间
	return dfs(l,k-1) && dfs(k+1,r)
}