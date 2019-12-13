package sorts

import "fmt"

// BubbleSort2 冒泡排序思想：
/*
	invalid operation: "冒泡排序2后：" + array (mismatched types string and []int)
*/
func BubbleSort2(array []int) {
	for i := 0; i < len(array); i++ {

		for j := 0; j < len(array)-1; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	fmt.Println("冒泡排序==2==后：", array)
}
