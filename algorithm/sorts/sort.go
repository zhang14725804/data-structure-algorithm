package sorts

import "fmt"

func SortMain() {
	origin := []int{5, 2, 7, 3, 6, 1}
	fmt.Println("排序之前的数据：", origin)
	BubbleSort(origin)
	result := QuickSort(origin)
	fmt.Println("快速排序之后：", result)
	InsertSort(origin)
	ShellSort(origin)
	SelectionSort(origin)
}
