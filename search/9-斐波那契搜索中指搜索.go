 

import "fmt"

func main_fab_search() {
	arr := make([]int, 1000, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = i
	}
	fmt.Println(fab_search(arr, 1117))
}

func makeFabArray(arr []int) []int {
	length := len(arr)
	fabLen := 2
	first, second, third := 1, 2, 3
	// 找出最接近的斐波那契
	for third < length {
		// 叠加计算斐波那契
		third, first, second = first+second, second, third
		fabLen++
	}
	fb := make([]int, fabLen)
	fb[0] = 1
	fb[1] = 1
	for i := 2; i < fabLen; i++ {
		fb[i] = fb[i-1] + fb[i-2]
	}
	return fb
}

/*
	斐波那契查找，空间换取时间
	todo：算法明白，实现难以理解，那还是不理解
*/
func fab_search(arr []int, target int) int {
	length := len(arr)
	// 匹配的斐波那契数组
	fab := makeFabArray(arr)
	fillLen := fab[len(fab)-1] // 填充长度
	fillArr := make([]int, fillLen)
	for i, v := range arr {
		fillArr[i] = v
	}
	lastdata := arr[length-1] //填充最后一个大数
	for i := length; i < fillLen; i++ {
		fillArr[i] = lastdata
	}
	left, mid, right := 0, 0, length
	index := len(fab) - 1 //游标
	for left < right {
		mid = left + fab[index-1] - 1
		if target < fillArr[mid] {
			right = mid - 1
			index--
		} else if target > fillArr[mid] {
			left = mid + 1
			index -= 2
		} else {
			// 越界
			if mid > right {
				return right
			}
			return mid
		}
	}
	return -1
}