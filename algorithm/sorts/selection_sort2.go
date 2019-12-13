package sorts

import "fmt"

// SelectionSort2 选择排序思想：
/*
	每次从余下的数中找最小的，并排到余下的数的最开头。
*/
func SelectionSort2(array []int) {
	length := len(array)
	if length < 2 {
		return
	}
	for i := 0; i < length; i++ {
		min := i
		for j := length - 1; j > i; j-- {
			if array[j] < array[min] {
				min = j
			}
		}
		array[min], array[i] = array[i], array[min]
	}
	fmt.Println("选择排序==2==之后：：", array)
}
