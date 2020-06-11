package main

import "fmt"

/*
	归并排序（和快排有些类似）
*/
func MergeSort(arr []int) []int {
	l := len(arr)
	/*
		todo:改良，len<10,改用插入排序
	*/
	if l <= 1 {
		return arr
	}
	mid := l / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	// 左右索引
	leftIndex := 0
	rightIndex := 0
	res := []int{}
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			res = append(res, left[leftIndex])
			leftIndex++
		} else if left[leftIndex] > right[rightIndex] {
			res = append(res, right[rightIndex])
			rightIndex++
		} else {
			res = append(res, left[leftIndex])
			res = append(res, right[rightIndex])
			leftIndex++
			rightIndex++
		}
	}
	// 余下的数插入
	for leftIndex < len(left) {
		res = append(res, left[leftIndex])
		leftIndex++
	}
	for rightIndex < len(right) {
		res = append(res, right[rightIndex])
		rightIndex++
	}
	return res
}

func main() {
	arr := []int{9, 2, 6, 1, 7, 0}
	fmt.Println(MergeSort(arr))
}
