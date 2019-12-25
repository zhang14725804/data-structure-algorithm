package sorts

import "fmt"

// CountSort 计数排序思想
/*
	（利用hash键值对，空间换取时间）：

	计数排序的核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。
	作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。
	{
		1:1,
		2:1,
		3:0,
		4:2,
		5:0,
		6:1,
	}
*/
func CountSort(array []int, maxValue int) {
	bucketLen := maxValue + 1
	// 初始为0的数组
	bucket := make([]int, bucketLen)

	sortedIndex := 0
	length := len(array)

	// 数组填充数据
	for i := 0; i < length; i++ {
		// [0 1 1 1 0 1 1 1]
		bucket[array[i]]++
	}
	fmt.Println("计数排序bucket：：", bucket)
	for j := 0; j < bucketLen; j++ {
		// 填充数据（包括重复数据）
		for bucket[j] > 0 {
			array[sortedIndex] = j
			sortedIndex++
			bucket[j]--
		}
	}
	fmt.Println("计数排序之后：：", array)
}
