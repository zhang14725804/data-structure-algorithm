package main

import "fmt"

/*
	选择排序（和冒泡排序有点相似）
	循环数组，每次从剩下的部分选择最大值\最小值，交换位置
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
	arr:=[]int{9,2,5,0,6,1,7,8,4,3}
	fmt.Println(SelectSort(arr))
}