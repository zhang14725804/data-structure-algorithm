package sorts

import "fmt"

/*
	计数排序思想（利用hash键值对，空间换取时间）：

	计数排序的核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。
	作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。
*/
func CountSort2(array []int, maxValue int) {
	bucketLen := maxValue+1
	// 申请存储空间
	bucket:=make([]int,bucketLen)

	for i:=0;i<len(array);i++{
		bucket[array[i]] +=1
	}
	fmt.Println("计数排序bucket：：", bucket)
	// array下角标
	sortIndex:=0
	for i:=0;i<len(bucket);i++{
		for bucket[i]>0{
			array[sortIndex] = i
			bucket[i]--
			sortIndex++
		}
	}
	fmt.Println("计数排序==2==之后：：", array)
}
