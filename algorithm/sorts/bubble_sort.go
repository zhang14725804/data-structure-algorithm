package sorts

import "fmt"

// BubbleSort 冒泡排序思想：
/*
	从尾部开始比较相邻的两个元素，如果尾部的元素比前面的大，就交换两个元素的位置。
	往前对每个相邻的元素都做这样的比较、交换操作，这样到数组头部时，第 1 个元素会成为最大的元素。
	重新从尾部开始第 1、2 步的操作，除了在这之前头部已经排好的元素。
	继续对越来越少的数据进行比较、交换操作，直到没有可比较的数据为止，排序完成。
*/
func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		//
		for j := 0; j < len(array)-1; j++ {
			if array[i] < array[j] {
				array[j], array[i] = array[i], array[j]
			}
		}
	}
	fmt.Println(array)
}
