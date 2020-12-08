/*
	打家劫舍 III
	这个地方的所有房屋的排列类似于一棵二叉树
*/

/*
	方法1：递归：自顶向下
*/
var hash map[*TreeNode]int = make(map[*TreeNode]int, 0)

func rob1(root *TreeNode) int {
	if root == nil {
		return 0
	}
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

/*
	方法2：递归：自顶向下
	无需cache缓存
*/
func rob(root *TreeNode) int {
	res := dfs(root)
	return MaxInt(res[0], res[1])
}

/*
	返回长度为2的数组
	arr[0]：不抢
	arr[1]：抢
*/
func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	// 抢，下家不能再抢
	rob := root.Val + left[0] + right[0]
	// 不抢，下家可抢可不抢，取决于收益
	not_rob := MaxInt(left[0], left[1]) + MaxInt(right[0], right[1])
	return []int{not_rob, rob}
}
