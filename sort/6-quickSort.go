 

import "fmt"

/*
	快速排序
*/
func QuickSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	} else {
		// base可以取随机一个数，也可以随机生成一个数
		base := arr[0]
		smaller := make([]int, 0, 0)
		bigger := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		mid = append(mid, base)
		for i := 1; i < l; i++ {
			if arr[i] < base {
				smaller = append(smaller, arr[i])
			} else if arr[i] > base {
				bigger = append(bigger, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		// 递归处理(todo:这里可以用栈实现)
		smaller, bigger = QuickSort(smaller), QuickSort(bigger)
		array := append(append(smaller, mid...), bigger...)
		return array
	}
}

// 改良版(todo:有重复的好像不行)
func QuickSort2(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	} else {
		//
		n := l - 1 // n>=0 && n <=l-1
		// base可以取随机一个数，也可以随机生成一个数
		base := arr[0]
		smaller := make([]int, 0, 0)
		bigger := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		mid = append(mid, base)
		for i := 1; i < l; i++ {
			if i == n {
				continue
			}
			if arr[i] < base {
				smaller = append(smaller, arr[i])
			} else if arr[i] > base {
				bigger = append(bigger, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		// 递归处理(todo:这里可以用栈实现)
		smaller, bigger = QuickSort(smaller), QuickSort(bigger)
		array := append(append(smaller, mid...), bigger...)
		return array
	}
}

func mainQuick() {
	arr := []int{9, 2, 5, 0, 6, 1, 7, 8, 4, 3}
	fmt.Println(QuickSort(arr))
	fmt.Println(QuickSort2(arr))
}

// 快速排序
func quick(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var left []int
	var right []int
	var mid []int
	for i := 0; i < len(arr); i++ {
		if arr[i] > pivot {
			right = append(right, arr[i])
		} else if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	return append(append(quick(left), mid...), quick(right)...)

}
