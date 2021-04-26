/*
	快速排序基本思想：
		通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序
*/
// 【递归】方式实现快排
func quickSort(nums []int) []int {
	// 😅 base case，递归出口
	if len(nums) < 2 {
		return nums
	}
	// 选择基准元素
	pivot := nums[0]
	var smaller, larger, res []int
	// 根据基准元素将数组分成左右两部分。😅😅😅注意此时nums取值范围nums[1:]，需要过滤pivot基准元素
	for _, num := range nums[1:] {
		if num < pivot {
			smaller = append(smaller, num)
		} else {
			larger = append(larger, num)
		}
	}
	// 😅 对左右两部分分别进行快排，然后拼接左边、右边、基准元素三部分，注意可变参数
	res = append(res, quickSort(smaller)...)
	res = append(res, pivot)
	res = append(res, quickSort(larger)...)
	return res
}

/*
	参考：https://www.geeksforgeeks.org/quick-sort/
	TODO:不会 😅😅😅
*/
func quickSort(nums []int, left, right int) []int {
	if left < right {
		pi := partition(nums, left, right)
		quickSort(nums, left, pi-1)
		quickSort(nums, pi+1, right)
	}
	return nums
}

/*
	This function takes last element as pivot, places the pivot element at its correct position in sortedarray,
	and places all smaller (smaller than pivot) to left of pivot and all greater elements to right of pivot
*/
func partition(nums []int, left, right int) int {
	// 基准元素
	pivot := nums[right]
	// 不懂 😅😅😅，为什么要left-1，而不是left
	i := left - 1
	for j := left; j <= right-1; j++ {
		if nums[j] < pivot {
			i++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	// 不懂 😅😅😅，这里又是干啥
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}
