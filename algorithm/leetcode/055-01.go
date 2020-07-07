/*
	给定一个非负整数数组，你最初位于数组的第一个位置。
	数组中的每个元素代表你在该位置可以跳跃的最大长度。
	判断你是否能够到达最后一个位置。

	思路1：动态规划(😅)
	思路2：贪心算法
*/
var dp []int
var nums []int
var length int
func canJump(arr []int) bool {
    nums = arr
	length = len(nums)
	// 初始化数组，长度为length，初始值为0（这么些是不是有点蠢）
	dp=make([]int,length)
	for i := 0; i < length; i++ {
		dp[i] = 0
	}
	dp[length-1] = 1
	return jump(0)
}
// 动态规划1 top-bottom
func jump(pos int) bool {
	if dp[pos] == 1{
		return true
	}else if dp[pos] == -1{
		return false
	}
	maxJump := compare(pos + nums[pos],length-1,false)
	for i := pos+1; i <= maxJump; i++ {
		if jump(i) == true{
			dp[pos] = 1
			return true
		}
	}
	dp[pos] = -1
	return false
}

func compare(a, b int, max bool) int {
	// max 是否返回最大值
	if a > b {
		if max == true {
			return a
		}
		return b
	}
	if max == true {
		return b
	}
	return a
}

// 动态规划2 bottom-top（倒着推）
