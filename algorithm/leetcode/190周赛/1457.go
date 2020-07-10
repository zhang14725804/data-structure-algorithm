/*
	给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。
	请你返回从 根到叶子节点 的所有路径中 伪回文 路径的数目。

	思路：伪回文等价于路径中最多只有一个点出现的次数是奇数
*/
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var hash = make(map[int]int,0)
func pseudoPalindromicPaths (root *TreeNode) int {
	// 存入hash
	hash[root.Val]++
	if root.Left == nil && root.Right == nil{
		t:=0
		for _,val:=range hash{
			// 某个点出现的次数是否是奇数
			if val % 2 == 1{
				t++
			}
		}
		// 根节点只能算一次，所以这里需要减
		hash[root.Val]--
		// 当前路径满足
		if t<=1{
            return 1
        } 
        return 0
	}
	res:=0
	// 遍历左右节点
	if root.Left!=nil{
		res+=pseudoPalindromicPaths(root.Left)
	}
	if root.Right!=nil{
		res+=pseudoPalindromicPaths(root.Right)
	}
	// 根节点只能算一次，所以这里需要减
	hash[root.Val]--
	return res
}