package main

import "fmt"

func MergeSort(arr []int) []int{
	l := len(arr)
	/*
		todo:改良，len<10,改用插入排序
	*/
	if l<= 1{
		return arr
	}else{
		mid := l/2
		left := MergeSort(arr[:mid])
		right := MergeSort(arr[mid:])
		return merge(left,right)
	}
}

func merge(left,right []int) []int  {
	// 左右索引
	leftIndex := 0
	rightIndex := 0
	array := []int{}
	for leftIndex <len(left) && rightIndex<len(right){
		if left[leftIndex] < right[rightIndex]{
			array = append(array,left[leftIndex])
			leftIndex++
		}else if left[leftIndex] > right[rightIndex]{
			array = append(array,right[rightIndex])
			rightIndex++
		}else{
			array = append(array,left[leftIndex])
			array = append(array,right[rightIndex])
			leftIndex++
			rightIndex++
		}
	}
	// 余下的数插入
	for leftIndex < len(left){
		array = append(array,left[leftIndex])
		leftIndex++
	}
	for rightIndex < len(right){
		array = append(array,right[rightIndex])
		rightIndex++
	}
	return array
}

func main()  {
	arr:=[]int{9,2,6,1,7,0}
	fmt.Println(MergeSort(arr))
}
