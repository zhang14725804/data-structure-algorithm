/*
	打家劫舍 III
	这个地方的所有房屋的排列类似于一棵二叉树
*/

/*
	方法1：递归：自顶向下
*/
var hash map[*TreeNode]int = make(map[*TreeNode]int, 0) // 备忘录

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 避免重叠子问题导致的重复计算
	if val, ok := hash[root]; ok {
		return val
	}
	// 抢，然后抢下下家
	doRob := root.Val
	if root.Left != nil {
		doRob += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		doRob += rob(root.Right.Left) + rob(root.Right.Right)
	}
	// 不抢，抢下一家
	donotRob := rob(root.Left) + rob(root.Right)

	res := MaxInt(doRob, donotRob)
	hash[root] = res
	return res
}