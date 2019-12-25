package sorts

import "fmt"

var sortMethods []string{
	"MergeSort",
	"ShellSort",
}
type Sort struct{
	result []int
}
func SortMain() {
	origin := []int{5, 2, 7, 3, 6, 1}

	// fmt.Println("排序之前的数据：", origin)
	// BubbleSort(origin)
	// BubbleSort2(origin)
	// fmt.Println("快速排序之后：", QuickSort(origin))
	// fmt.Println("快速排序==2==之后：", QuickSort2(origin))
	// InsertSort(origin)
	// InsertSort2(origin)
	// CountSort(origin, 7)
	// CountSort2(origin, 7)
	// SelectionSort(origin)
	// SelectionSort2(origin)
	// ShellSort(origin)
	// fmt.Println("归并排序之后：", MergeSort(origin))
	// // 四种排序不会
	// BucketSort(origin)
	// RadixSort(origin)
	// HeapSort(origin)
}
