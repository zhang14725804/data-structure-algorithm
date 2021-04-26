/*
	快速排序基本思想：
		通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序
*/
func quickSort(nums []int) []int {
	// 😅 base case，递归出口
	if len(nums) < 2 {
		return nums
	}
	// 选择基准元素
	pivot := nums[0]
	var smaller, larger, res []int
	// 根据基准元素将数组分成左右两部分
	for _, num := range nums {
		if num > pivot {
			smaller = append(smaller, num)
		} else {
			larger = append(larger, num)
		}
	}
	// 😅 对左右两部分分别进行快排，然后拼接左边、右边、基准元素三部分
	res := append(res, quickSort(smaller...))
	res := append(res, pivot)
	res := append(res, quickSort(larger...))
	return res
}