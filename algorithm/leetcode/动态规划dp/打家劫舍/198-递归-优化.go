/*
	递归（自顶向下）
	解决重叠子问题
*/

var nums []int
var memo []int // 备忘录

func rob(_nums []int) int {
	nums = _nums
	memo = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = -1
	}

	return dp(0)
}
func dp(start int) int {
	// base case
	if start >= len(nums) {
		return 0
	}
	// 避免重叠子问题导致的重复计算
	if memo[start] != -1 {
		return memo[start]
	}
	// 对于第start家。
	// dp(start+1)不抢，抢下一家；
	// nums[start]+dp(start+2)抢这一家，再抢下下一家
	res := MaxInt(dp(start+1), nums[start]+dp(start+2))
	memo[start] = res
	return res
}