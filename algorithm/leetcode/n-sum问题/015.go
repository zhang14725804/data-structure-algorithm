/*
	思路：
	1. 先将数组排序；
	2. 三个指针，一个i，一个start，一个end。每次循环i，然后i固定，start等于i+1，end等于最后一个数；
	（1）每次相加i，start，end，如果大于target，end向前移动
	（2）如果小于target，start向前移动；
	（3）如果相等，i，start，end加入结果集
	（4）遇到相同的数字跳过，i，start，end都有可能遇到相等的数字
*/
func threeSum(nums []int) [][]int {
	n := len(nums)
	// 排序先
	nums = sort(nums)

	var res [][]int
	for i := 0; i < n; i++ {
		// 排除重复的情况
		if i == 0 || nums[i] != nums[i-1] {
			// 双指针
			start, end := i+1, n-1
			// 收缩start，end的情况
			for start < end {
				// 符合条件
				if nums[i]+nums[start]+nums[end] == 0 {
					res = append(res, []int{nums[i], nums[start], nums[end]})
					// 0102 把这一步忘了
					start++
					end--
					// 排除重复的情况
					// （0）for 而不是【if】
					// （1）start < end
					// （2）上一个数和当前数相等
					for start < end && nums[start] == nums[start-1] {
						start++
					}
					for start < end && nums[end] == nums[end+1] {
						end--
					}
				} else if nums[i]+nums[start]+nums[end] < 0 {
					start++
				} else {
					end--
				}
			}
		}
	}
	return res
}