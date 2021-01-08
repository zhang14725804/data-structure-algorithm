 

import "fmt"

/*
	插入排序
	从左往右，从右往左两种
*/
func InsertSort1(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	}
	// 跳过第一个
	for i := 1; i < l; i++ {
		backup := arr[i] // 备份插入的数据
		j := i - 1       // 上一个位置循环找到位置插入
		for j >= 0 && backup < arr[j] {
			arr[j+1] = arr[j] // 从前往后移动
			j--
		}
		// 插入
		arr[j+1] = backup
	}
	return arr
}

func InsertSort2(array []int) []int {
	length := len(array)
	if length < 2 {
		return array
	}
	for i := 1; i < length; i++ {
		// 注意：j >= 0
		for j := i - 1; j >= 0; j-- {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func mainInsert() {
	arr := []int{9, 2, 5, 0, 6, 1, 7, 8, 4, 3}
	fmt.Println(InsertSort1(arr))
}
