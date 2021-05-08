/*
	给定一个非负整数数组，你最初位于数组的第一个位置。
	数组中的每个元素代表你在该位置可以跳跃的最大长度。
	判断你是否能够到达最后一个位置。
*/
/*
	思路：贪心算法 question
	通过题目中的跳跃规则，最多能跳多远？如果能够越过最后一格，返回 true，否则返回 false。
*/
func canJump(nums []int) bool {
	n := len(nums)
	farthest := 0
	for i := 0; i < n-1; i++ {
		// 计算能跳的最远距离
		farthest = MaxInt(farthest, i+nums[i])
		// 遇到0
		if farthest <= i {
			return false
		}
	}

	return farthest >= n-1
}
