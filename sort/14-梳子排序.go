/*
	梳子排序：希尔排序的最优版
*/

func ComoSort(arr []int) []int {
	l := len(arr)
	gap := l
	for gap > 1 {
		// 为什么这么做
		gap = gap * 10 / 13
		// 收缩步长交换位置
		for i := 0; i+gap < l; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
			}
		}
	}
	return arr
}

func mainComoSort() {
	array := []int{16, 8, 1, 24, 30}
	fmt.Println(ComoSort(array))
}
