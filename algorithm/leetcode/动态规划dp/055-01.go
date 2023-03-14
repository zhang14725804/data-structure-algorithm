/*
	question
	😅😅😅
	思路1：动态规划(😅)
	dp每个位置有三种可能
		1：可到达
		0：未知
		-1：不可到达
*/
var dp []int
var nums []int
var length int

func canJump(arr []int) bool {
	nums = arr
	length = len(nums)
	// 初始化数组，长度为length，初始值为0（这么些是不是有点蠢）
	dp = make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 0
	}
	dp[length-1] = 1
	// 动态规划1 top->bottom 从第一个元素开始递归
	return jump(0)
}
func jump(pos int) bool {
	if dp[pos] == 1 {
		return true
	} else if dp[pos] == -1 {
		return false
	}
	maxJump := MinInt(pos+nums[pos], length-1)
	for i := pos + 1; i <= maxJump; i++ {
		if jump(i) == true {
			dp[pos] = 1
			return true
		}
	}
	dp[pos] = -1
	return false
}