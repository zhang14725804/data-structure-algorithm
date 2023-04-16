/*
	方法1：双指针算法，快慢指针
		1. slow只走不重复的数据
		2. fast每次都走
		3. 遇到重复的数据，slow跳过
		4. 遇到不重复的数据，slow先前进一步，并更新nums[slow]的值
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// slow存所有不重复的数
	slow, fast := 0, 0
	// fast从前向后遍历
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			// 😅😅😅 slow需要先走，跳过【0】号元素，第一个元素不存在重复问题
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}