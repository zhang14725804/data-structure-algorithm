 

import "fmt"

/*
	希尔排序（多线程排序）
	不断地收缩步长
*/
func ShellSort(arr []int) []int {
	l := len(arr)
	if l <= 1 {
		return arr
	} else {
		gap := l / 2
		for gap > 0 {
			//
			for i := 0; i < gap; i++ {
				step(arr, i, gap)
			}
			gap /= 2 // gap--
		}
	}
	return arr
}

// 每个步长在进行插入排序
func step(arr []int, start, gap int) {
	l := len(arr)
	for i := start + gap; i < l; i += gap {
		// todo：不太懂
		backup := arr[i] // 备份插入的数据
		j := i - gap     // 上一个位置循环找到位置插入
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j] // 从前往后移动
			j -= gap
		}
		// 插入
		arr[j+gap] = backup
	}
}
func main() {
	arr := []int{9, 2, 5, 0, 6, 1, 7, 8, 4, 3}
	fmt.Println(ShellSort(arr))
}
