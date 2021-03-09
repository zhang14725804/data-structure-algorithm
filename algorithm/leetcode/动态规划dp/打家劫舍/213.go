/*
	打家劫舍2
	这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。
*/

/*
	可能有三种不同情况：要么都不被抢；要么第一间房子被抢最后一间不抢；要么最后一间房子被抢第一间不抢
	只要比较情况二和情况三就行了，因为这两种情况对于房子的选择余地比情况一大
*/

var nums []int

func rob(_nums []int) int {
	nums = _nums
	size := len(nums)
	// 临界条件
	if size == 0 {
		return 0
	}
	if size == 1 {
		return nums[0]
	}
	// 是否考虑第一家，分为两种情况
	// 抢第一家：helper(0, size-1)
	// 不抢第一家：helper(1, size)
	return MaxInt(helper(0, size-1), helper(1, size))
}

func helper(start, end int) int {
	// base case
	pre := 0
	cur := nums[start]
	// 状态转移只和dp[i]最近的两个状态有关
	for i := start + 2; i <= end; i++ {
		temp := cur
		cur = MaxInt(cur, nums[i-1]+pre)
		pre = temp
	}
	return cur
}