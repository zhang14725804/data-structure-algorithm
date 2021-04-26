/*
	冒泡排序基本思想：
		它重复地走访过要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。
		走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。
		这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。
*/
func BubbleSort(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	// 😅😅😅 只剩一个不需要冒泡
	for i := 0; i < l; i++ {
		// 算法优化
		needChange := false
		// 注意j取值范围，😅😅😅
		for j := 0; j < l-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				needChange = true
			}
		}
		if !needChange {
			break
		}
	}
	return nums
}
