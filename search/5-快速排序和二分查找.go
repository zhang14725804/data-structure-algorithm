package main

import "fmt"

/*
	快速排序
*/
func main(){
	arr:=[]int{5,4,6,7,4,8,3,2,9,0,1}
	fmt.Println(QuickSort(arr))
	fmt.Println(binSearch(arr,12))
}

// 快速排序
func QuickSort(arr []int)[]int{
	l:=len(arr)
	if l<=1{
		return arr
	}else{
		base := arr[0]
		small := make([]int,0,0)
		bigger := make([]int,0,0)
		mid := make([]int,0,0)
		//
		mid = append(mid,base)
		for i:=1; i<l; i++{
			if arr[i]<base{
				small = append(small,arr[i])
			}else if arr[i]>base{
				bigger = append(bigger,arr[i])
			}else{
				mid = append(mid,arr[i])
			}
		}
		//递归
		small,bigger = QuickSort(small),QuickSort(bigger)
		return append(append(small,mid...),bigger...)
	}
}

// 二分查找
func binSearch(arr []int,target int) int  {
	fmt.Println(arr,target)
	left:=0
	right:=len(arr)-1
	for left<right{
		mid:=(left+right)/2
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