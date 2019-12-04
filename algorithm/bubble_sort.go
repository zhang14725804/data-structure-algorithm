package algorithm

import "fmt"

/*
	冒泡排序
*/
func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		//
		for j := 0; j < len(array)-1; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println(array)
}
