 

import "fmt"

// 找到极大值
func selectMax(arr []int) int {
	l := len(arr)
	if l <= 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < l; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}

// 对每个数量级排序
func bitSort(arr []int, bit int) []int {
	l := len(arr)
	counts := make([]int, 10) // 统计长度
	for i := 0; i < l; i++ {
		num := (arr[i] / bit) % 10 // 分层处理（区分数量级）
		counts[num]++              // 统计余数相等的个数
	}
	fmt.Println(counts)
	// todo:不懂了
	for i := 1; i < 10; i++ {
		counts[i] += counts[i-1]
	}
	fmt.Println(counts)
	temp := make([]int, 10)
	for i := l - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		temp[counts[num]-1] = arr[i] // 计算排序的位置
		counts[num]--
	}
	for i := 0; i < l; i++ {
		arr[i] = temp[i] // 保存数组
	}
	return arr
}

/**
基数排序（todo：有点绕）
*/
func RadixSort(arr []int) []int {
	max := selectMax(arr) // 寻找最大值
	// 按照数量级分段
	for bit := 1; max/bit > 0; bit *= 10 {
		arr = bitSort(arr, bit)
	}
	return arr
}
func main() {
	arr := []int{99, 21, 585, 305, 6111, 1142, 79, 834, 4298, 3097}
	fmt.Println(RadixSort(arr))
}