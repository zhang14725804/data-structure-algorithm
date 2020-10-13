/*
	给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
	返回滑动窗口中的最大值。

	进阶：你能在线性时间复杂度内解决此题吗？

	单调队列问题（滑动窗口）
*/

// 解法3：单调队列（todo，不懂）
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

// 解法一：暴力破解（超时）
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	res := make([]int, n-k+1)
	// 循环条件
	for i := 0; i < len(res); i++ {
		max := INT_MIN
		for j := 0; j < k; j++ {
			max = MaxInt(max, nums[j+i])
		}
		res[i] = max
	}
	return res
}

// 解法2：优先队列（todo）