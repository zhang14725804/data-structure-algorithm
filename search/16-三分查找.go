package main

import "fmt"

func main(){
	arr := make([]int,1000,1000)
	for i:=0; i<1000;i++  {
		arr[i] = i
	}
	fmt.Println(threeSearch(arr,29))
}

/*
	三分查找
*/
func threeSearch(arr []int,target int)  int {
	left := 0
	right := len(arr)-1
	i:=0
	// <=
	for left <= right {
		i++
		fmt.Println("查找次数:",i)
		mid1 := left + int((right-left)/3)
		mid2 := right - int((right-left)/3)
		midData1 := arr[mid1]
		midData2 := arr[mid2]
		if midData1 == target{
			return mid1
		} else if midData2 == target{
			return mid2
		}
		//
		if midData1 < target {
			left = mid1+1
		}else if midData2 > target{
			right = mid2-1
		}else{
			right = right - 1
			left = left + 1
		}
	}
	return -1
}