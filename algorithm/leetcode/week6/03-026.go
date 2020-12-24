/*
	给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
	不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/

/*
	方法1：双指针算法，快慢指针
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// slow存所有不重复的数
	slow, fast := 0, 0
	// fast从前向后遍历
	for fast < len(nums) {
		// 判断是否和上一个数相同
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}