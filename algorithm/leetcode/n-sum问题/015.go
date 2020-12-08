/*
	给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

	注意：答案中不可以包含重复的三元组。
*/

/*
	思路：
	先将数组排序；
	三个指针，一个i，一个start，一个end，遍历数组，i从0开始，start等于i+1，end等于最后一个数；
	每次相加i，start，end，如果大于target，end向前移动，如果小于target，start向前移动；
	遇到i和start相等，跳过
*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	// 对数组排序(冒泡，蠢一些)
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}

	numsLen := len(nums)
	for i := 0; i < numsLen-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			start, end := i+1, numsLen-1
			for start < end {
				if nums[i]+nums[start]+nums[end] == 0 {
					res = append(res, []int{nums[i], nums[start], nums[end]})
					start++
					end--
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
