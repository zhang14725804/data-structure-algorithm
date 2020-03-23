package main

import "fmt"

/*
	奇偶排序(多线程？)
	适用于类似对身高排序
*/
func OddEvenSort(arr []int) []int {

	isSorted := false
	for ;isSorted ==false;{
		isSorted = true
		// 奇数位
		for i:=1;i<len(arr)-1 ;i=i+2  {
			if arr[i]>arr[i+1]{
				arr[i],arr[i+1] = arr[i+1],arr[i]
				isSorted = false
			}
		}
		// 偶数位
		for i:=0;i<len(arr)-1 ;i=i+2  {
			if arr[i]>arr[i+1]{
				arr[i],arr[i+1] = arr[i+1],arr[i]
				isSorted = false
			}
		}
	}
	return arr
}

func main()  {
	arr:=[]int{9,2,6,1,7,0}
	fmt.Println(OddEvenSort(arr))
}
