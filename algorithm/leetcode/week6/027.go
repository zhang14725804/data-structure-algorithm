/*
	方法1：快慢指针,同时向右走
*/
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

/*
	方法2：双指针2，相向而行
*/
func removeElement(nums []int, val int) int {
	i, n := 0, len(nums)
	for i < n {
		// 找到目标元素
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}
	return n
}