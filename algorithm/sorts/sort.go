package sorts

import (
	"fmt"
	"reflect"
)

// Sort 排序结构
// type Sort struct {
// 	Result []int
// }
type T struct{}

// SortMain 排序入口函数
func SortMain() {
	sortMethods := []string{
		"bar",
		"foo",
	}
	t := &T{}
	// origin := []int{5, 2, 7, 3, 6, 1}
	for _, s := range sortMethods {
		reflect.ValueOf(t).MethodByName(s).Call(nil)
	}

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

func bar() {
	fmt.Println("======bar======")
}
func foo() {
	fmt.Println("======foo======")
}
