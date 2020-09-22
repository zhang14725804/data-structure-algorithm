/*
	打家劫舍2

	这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。
*/
var nums []int
var n int

func rob(_nums []int) int {
	nums = _nums
	n = len(nums)
	// 是否考虑第一家，分为两种情况
	return MaxInt(helper(0, n-1), helper(1, n))
}

func helper(start, end int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	pre := 0
	cur := nums[start]
	// pre,cur
	for i := start + 2; i <= end; i++ {
		temp := cur
		cur = MaxInt(pre+nums[i-1], cur)
		pre = temp
	}
	return cur
}