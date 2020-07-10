/*
	给你两个数组 nums1 和 nums2 。
	请你返回 nums1 和 nums2 中两个长度相同的 非空 子序列的最大点积。
	数组的非空子序列是通过删除原数组中某些元素（可能一个也不删除）后剩余数字组成的序列，
	但不能改变数字间相对顺序。比方说，[2,3,5] 是 [1,2,3,4,5] 的一个子序列而 [1,5,3] 不是。

	思路：动态规划
*/
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX
func maxDotProduct(nums1 []int, nums2 []int) int {
	n,m:=len(nums1),len(nums2)
	// 初始化二维数组并赋值
	dp:=make([][]int,n+1)
	for i := 0; i < n+1; i++ {
		dp[i]=make([]int,m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = INT_MIN
		}
	}
	// 第一行第一列都是0
	for i := 0; i < n+1; i++ {
		dp[i][0] = 0
	}
	for i := 0; i < m+1; i++ {
		dp[0][i] = 0
	}
	// todo：动态规划还是难以理解
	res:=INT_MIN
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = compare(dp[i-1][j],dp[i][j-1],true)
			t:=dp[i-1][j-1] + nums1[i-1]*nums2[j-1]
			res = compare(res,t,true)
			dp[i][j] = compare(dp[i][j],t,true) 
		}
	}
	return res
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