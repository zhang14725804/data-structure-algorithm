package main

import "fmt"

/*
	二分查找改良版--斐波那契查找
	斐波那契方式切割
*/
func binSearchMid(arr []int,target int) int  {
	fmt.Println(arr,target)
	left:=0
	right:=len(arr)-1

	i:=0
	for left<=right{
		i++
		fmt.Println("查找次数：",i)

		// mid:=(left+right)/2
		leftv := float64(target-arr[left]) // 大段
		allv := float64(arr[right]-arr[left]) // 整段
		diff := float64(right-left)
		mid := int(float64(left)+leftv/allv*diff) // 中间值

		if mid<0 || mid>=len(arr){
			return -1
		}
		if arr[mid]>target{
			right = mid-1
		}else if arr[mid]<target{
			left = mid+1
		}else{
			return mid
		}
	}
	return -1
}

func mainBinSearchMid()  {
	arr := make([]int,1000,1000)
	for i:=0; i<1000;i++  {
		arr[i] = i
	}
	binSearchMid(arr,4)
}