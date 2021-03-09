/*
	打家劫舍 III
	这个地方的所有房屋的排列类似于一棵二叉树
*/

/*
	方法1：递归：自顶向下
*/
var memo map[*TreeNode]int = make(map[*TreeNode]int, 0)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 避免重叠子问题导致的重复计算
	if val, ok := memo[root]; ok {
		return val
	}
	// 抢，抢下下家
	doRob := root.Val
	if root.Left != nil {
		doRob += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		doRob += rob(root.Right.Left) + rob(root.Right.Right)
	}
	// 不抢，抢下家
	donotRob := rob(root.Left) + rob(root.Right)
	// 取两者较大值
	memo[root] = MaxInt(doRob, donotRob)
	return memo[root]
}