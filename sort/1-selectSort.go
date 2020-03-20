package main

import "fmt"

/*
	选择排序（和冒泡排序有点相似）
*/
func SelectSort(arr []int) []int{
	l := len(arr)
	if l <=1 {
		return arr
	}

	// 只剩下第一个元素，不需要挑选
	for i:= 0; i<l-1;i++{
		min := i // 标记索引
		// 每次选出一个极小值
		for j:=i+1;j<l;j++{
			if arr[min]<arr[j]{
				min = j // 保存极小值索引
			}
		}
		if min != i{
			arr[i],arr[min] = arr[min],arr[i]
		}
	}
	return arr
}
func mainInt()  {
	arr:=[]int{4,7,2,1,3,8,5,9,6,8}
	fmt.Println(SelectSort(arr))
}