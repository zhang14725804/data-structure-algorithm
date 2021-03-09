/*
	方法2：递归：自顶向下
	无需cache缓存(question 妙啊)
*/
func rob(root *TreeNode) int {
	res := helper(root)
	return MaxInt(res[0], res[1])
}

/*
	返回长度为2的数组
	arr[0]：不抢
	arr[1]：抢
*/
func helper(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := helper(root.Left)
	right := helper(root.Right)
	// 抢，下家不能再抢
	doRob := root.Val + left[0] + right[0]
	// 不抢，下家可抢可不抢，取决于收益（机智👉 🐔）
	donotRob := MaxInt(left[0], left[1]) + MaxInt(right[0], right[1])
	return []int{donotRob, doRob}
}