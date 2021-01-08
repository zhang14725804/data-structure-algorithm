 

import "fmt"

/*
	找到第K大的数:
	（1）堆排序
	（2）快速排序的修正算法
*/
func mainTheK() {
	arr := []int{9, 7, 8, 6, 4, 5, 1, 3, 2, 0}
	fmt.Println(findTheK(arr, 2))
}

func findTheK(arr []int, k int) int {
	return find(arr, 0, len(arr)-1, k)
}
func find(arr []int, left, right, k int) int {
	if left >= right {
		return arr[left]
	}
	query := partition(arr, left, right)
	if query+1 == k {
		//第N大的数
		return arr[query]
	}
	if k < query+1 {
		return find(arr, left, query-1, k)
	}
	return find(arr, left, query+1, k)
}

// 分区函数，根据基准元素将数组分成两部分，返回基准元素下表（以递增排序为例）
func partition(arr []int, left, right int) int {
	// 将最后一个元素作为分区标识
	pivot := right
	i := left
	for j := left; j < pivot; j++ {
		//
		if arr[j] > arr[pivot] {
			swap(arr, i, j)
			i++
		}
	}
	swap(arr, i, pivot)
	return i
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}