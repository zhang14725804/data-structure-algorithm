/*
	给定一个非负整数数组，你最初位于数组的第一个位置。
	数组中的每个元素代表你在该位置可以跳跃的最大长度。
	判断你是否能够到达最后一个位置。

	思路1：动态规划(😅)
	思路2：贪心算法
*/
// 贪心算法
func canJump(nums []int) bool {
	maxJump:=len(nums)-1
	for i := len(nums)-2; i >=0 ; i-- {
		// 这骚操作
		if i+nums[i]>=maxJump{
			maxJump = i
		}
	}
    return maxJump == 0
}
