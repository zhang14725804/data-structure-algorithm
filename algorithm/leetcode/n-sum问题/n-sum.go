/*
	n-sum问题，比如100sum
	排序数组（careful）
	答案中不可以包含重复的元组。

	递归的方式
*/

func nSumTarget(nums []int, target, start, n int) [][]int {
	var res [][]int
	numsLen := len(nums)
	// 至少是2sum，且数组大小不小于n
	if n < 2 || numsLen < n {
		return res
	}
	// twoSum
	if n == 2 {
		left, right := start, numsLen-1
		for left < right {
			if nums[left]+nums[right] == target {
				res = append(res, []int{nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}
	} else {
		for i := start; i < numsLen; i++ {
			// 递归
			sub := nSumTarget(nums, target-nums[i], i+1, n-1)
			for _, arr := range sub {
				arr = append(arr, nums[i])
				res = append(res, arr)
			}
			for i < numsLen-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}