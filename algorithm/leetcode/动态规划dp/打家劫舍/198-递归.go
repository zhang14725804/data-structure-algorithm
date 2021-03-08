/*

	你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
	如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
	给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

*/

/*
	递归（自顶向下）
	递归中会计算很多重复的解（重叠子问题）
*/
var nums []int

func rob(_nums []int) int {
	nums = _nums
	return dp(0)
}
func dp(start int) int {
	// base case
	if start >= len(nums) {
		return 0
	}
	// 对于第start家。
	// dp(start+1)不抢，抢下一家；
	// nums[start]+dp(start+2)抢这一家，再抢下下一家
	return MaxInt(dp(start+1), nums[start]+dp(start+2))
}