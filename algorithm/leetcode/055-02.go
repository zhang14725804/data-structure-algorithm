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
	/*
		dp每个位置有三种可能
		1：可到达
		0：未知
		-1：不可到达
	*/
	// 动态规划2 bottom-up 从第最后个元素开始倒着递归
	// length-2最后一个数不需要考虑
	for i := length-2; i>=0; i-- {
		maxJump:=compare(i+nums[i],length-1,false)
		for j := i+1; j <= maxJump; j++ {
			if dp[j] == 1{
				dp[i] = 1
				break
			}
		}
	}
	if dp[0] == 1{
		return true
	}else{
		return false
	}
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