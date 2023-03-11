/*
	解法3：单调队列
	（todo，不懂）
*/
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	q := &Queue{}

	for i := 0; i < len(nums); i++ {
		if q.size() > 0 && i-k+1 > q.front() {
			q.pop_front()
		}
		for q.size() > 0 && nums[q.back()] <= nums[i] {
			q.pop_back()
		}
		q.push_back(i)
		if i >= k-1 {
			res = append(res, nums[q.front()])
		}
	}
	return res
}

/*
	解法一：暴力破解（超时）
	😅😅😅
*/
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	// (1) 声明长度【n-k+1】 而不是 0
	res := make([]int, n-k+1)
	// （2）循环条件【len(res)】 而不是 n
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
	解法2：优先队列（todo）
*/ 