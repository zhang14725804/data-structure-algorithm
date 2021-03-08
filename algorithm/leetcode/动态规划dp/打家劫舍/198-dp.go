/*
	自底向上
*/
func rob(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	dp := make([]int, size+1)
	// dp[i]=x 从第 i 间房子开始抢劫，最多能抢到的钱为 x
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= size; i++ {
		// 对于第i家。
		// dp(i-1)不抢，抢下一家；
		// nums[i-1]+dp(i-2)抢这一家，再抢下下一家
		dp[i] = MaxInt(dp[i-1], nums[i-1]+dp[i-2])
	}
	return dp[size]
}