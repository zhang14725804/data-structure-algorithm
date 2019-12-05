package sorts

import "fmt"

/*
	选择排序思想：

	每次从余下的数中找最小的，并排到余下的数的最开头。
*/
func SelectionSort(array []int) {
	length := len(array)
	if length < 2 {
		return
	}

	for i := 0; i < length; i++ {
		// 初始的最小值位置从0开始，依次向右
		min := i
		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := length - 1; j > i; j-- {
			if array[j] < array[min] {
				min = j
			}
		}
		// 把每次找出来的最小值与之前的最小值做交换
		array[i], array[min] = array[min], array[i]
	}
	fmt.Println("选择排序之后：：", array)
}
