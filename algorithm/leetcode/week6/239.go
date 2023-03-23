/*
	解法一：暴力破解（超时）😅
*/
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	// (1) 声明长度【n-k+1】 而不是 0
	res := make([]int, n-k+1)
	// （2）😅 循环条件【len(res)】 而不是 n， 避免再判断边界
	for i := 0; i < len(res); i++ {
		max := INT_MIN
		// 😅取当前窗口的最大值
		for j := 0; j < k; j++ {
			max = MaxInt(max, nums[j+i])
		}
		// (3) 更新res，直接更新对应位置的值，而不是append
		res[i] = max
	}
	return res
}

/*
	再次遇到还是不会 😅😅😅
	TODO
	优先队列
	单调队列
	分块+预处理
*/
