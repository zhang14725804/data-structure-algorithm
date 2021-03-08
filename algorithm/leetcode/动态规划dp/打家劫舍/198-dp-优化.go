/*
	动态规划优化（状态压缩） (question)
	状态转移只和dp[i]最近的两个状态有关
*/

func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	pre := 0
	cur := nums[0]
	for i := 2; i <= size; i++ {
		temp := cur
		// 对于第i家。
		// dp(i-1)不抢，抢下一家；
		// nums[i-1]+dp(i-2)抢这一家，再抢下下一家
		cur = MaxInt(cur, nums[i-1]+pre)
		pre = temp
	}
	return cur
}