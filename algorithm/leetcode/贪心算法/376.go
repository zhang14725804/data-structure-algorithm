func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// 前一对差值、当前一对差值
	preDiff, curDiff := 0, 0
	// 记录峰值个数，序列默认序列最右边有一个峰值
	res := 1
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if (curDiff > 0 && preDiff <= 0) || (preDiff >= 0 && curDiff < 0) {
			res++
			preDiff = curDiff
		}
	}
	return res
}