/*
	给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

	如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
*/

/*
	方法1：暴力破解
	从第 0 个数字开始，依次添加数字，记录当总和大于等于 s 时的长度。

	从第 1 个数字开始，依次添加数字，记录当总和大于等于 s 时的长度。

	从第 2 个数字开始，依次添加数字，记录当总和大于等于 s 时的长度。

	...

	从最后个数字开始，依次添加数字，记录当总和大于等于 s 时的长度。

	从上边得到的长度中选择最小的即可
*/
func minSubArrayLen(s int, nums []int) int {
	min := INT_MAX
	n := len(nums)

	for i := 0; i < n; i++ {
		start := i
		sum := 0
		for start < n {
			sum += nums[start]
			start++
			if sum >= s {
				min = compare(min, start-i, false)
				break
			}
		}
	}
	if min == INT_MAX {
		return 0
	}
	return min
}

/*
	方法2：双指针
*/
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left := 0
	right := 0
	sum := 0
	min := INT_MAX
	for right < n {
		sum += nums[right]
		right++
		for sum >= s {
			min := compare(min, right-left, false)
			sum -= nums[left]
			left++
		}
	}
	if min == INT_MAX {
		return 0
	}
	return min
}

/*
	方法三：二分法
*/