/*
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

	必须在原数组上操作，不能拷贝额外的数组。
	尽量减少操作次数。
*/

/*
	思路：双指针，快慢指针
	将所有非零元素放到开头，这样就保证末尾剩下的都是 0
	双指针，指针 fast 用于遍历数组，指针 slow 开始指向开头，保证它前边的所有元素都是非 0 元素。
*/
func moveZeroes(nums []int) {
	fast, slow := 0, 0
	// 当 fast 指针遇到非零元素就和 slow 指针指向的元素交换，slow 指针然后后移
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
}