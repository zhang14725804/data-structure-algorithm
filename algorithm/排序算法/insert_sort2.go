package sorts

import "fmt"

// InsertSort2 插入排序思想：
/*
	每一步将一个待排序的记录，插入到前面已经排好序的有序序列中去，直到插完所有元素为止。
*/
func InsertSort2(array []int) {
	length := len(array)
	if length < 2 {
		return
	}
	// i:=1，j>=0
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println("插入排序==2==之后：：", array)
}
