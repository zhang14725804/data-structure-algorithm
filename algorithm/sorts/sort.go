package sorts

import "fmt"

func SortMain() {
	origin := []int{5, 2, 7, 3, 6, 1}
	fmt.Println("排序之前的数据：", origin)
	BubbleSort(origin)
	BubbleSort2(origin)
	fmt.Println("快速排序之后：", QuickSort(origin))
	fmt.Println("快速排序==2==之后：", result)
	InsertSort(origin)
	InsertSort2(origin)
	CountSort(origin, 7)
	CountSort2(origin, 7)
	SelectionSort(origin)
	SelectionSort2(origin)

	// 五种排序不会
	MergeSort(origin)
	BucketSort(origin)
	RadixSort(origin)
	ShellSort(origin)
	HeapSort(origin)
}
