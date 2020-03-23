package main

import "fmt"

func QuickSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}else{
		// base可以取随机一个数，也可以随机生成一个数
		base := arr[0]
		smaller := make([]int,0,0)
		bigger := make([]int,0,0)
		mid := make([]int ,0,0)
		mid = append(mid,base)
		for i:=1;i<l;i++{
			if arr[i] < base{
				smaller = append(smaller,arr[i])
			} else if arr[i]>base{
				bigger = append(bigger,arr[i])
			}else{
				mid = append(mid,arr[i])
			}
		}
		// 递归处理(todo:这里可以用栈实现)
		smaller,bigger = QuickSort(smaller),QuickSort(bigger)
		array := append(append(smaller,mid...),bigger...)
		return array
	}
}

// 改良版(todo:有重复的好像不行)
func QuickSort2(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}else{
		//
		n:=l-1 // n>=0 && n <=l-1
		// base可以取随机一个数，也可以随机生成一个数
		base := arr[0]
		smaller := make([]int,0,0)
		bigger := make([]int,0,0)
		mid := make([]int ,0,0)
		mid = append(mid,base)
		for i:=1;i<l;i++{
			if i == n {
				continue
			}
			if arr[i] < base{
				smaller = append(smaller,arr[i])
			} else if arr[i]>base{
				bigger = append(bigger,arr[i])
			}else{
				mid = append(mid,arr[i])
			}
		}
		// 递归处理(todo:这里可以用栈实现)
		smaller,bigger = QuickSort(smaller),QuickSort(bigger)
		array := append(append(smaller,mid...),bigger...)
		return array
	}
}

func mainQuick()  {
	arr:=[]int{4,7,2,1,3,8,5,9,6,8}
	fmt.Println(QuickSort(arr))
	fmt.Println(QuickSort2(arr))
}
