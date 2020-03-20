package main

import "fmt"

/*
	冒泡排序
	每次循环确定一个数子的位置
*/

func BubbleSort(arr []int) []int{
	l:=len(arr)
	if l<=1{
		return arr
	}
	// 只剩一个不需要冒泡
	for i:=0;i<l-1;i++{
		// 算法优化
		needChange := false
		for j:=0;j<l-i-1;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
				needChange = true
			}
		}
		if !needChange{
			break
		}
	}
	return arr
}

func main()  {
	arr:=[]int{4,7,2,1,3,8,5,9,6,8}
	fmt.Println(BubbleSort(arr))
}