 

import "fmt"

/*
	todo:不好理解

	A找到第一个i
	B找到最后一个i
	C找到第一个大于等于i的
	D找到最后一个小于等于i的
*/
func main_search_n() {
	arr := []int{0, 1, 2, 2, 3, 3, 3, 4, 4, 5, 6, 6, 6, 7, 8, 8, 8, 8, 9}
	fmt.Println(bin_search_midA(arr, 3))
	fmt.Println(bin_search_midB(arr, 3))
	fmt.Println(bin_search_midC(arr, 5))
	fmt.Println(bin_search_midD(arr, 5))
}

// 第一个i
func bin_search_midA(arr []int, target int) int {
	fmt.Println(arr, target)
	left := 0
	right := len(arr) - 1
	index := -1
	i := 0
	for left <= right {
		i++
		fmt.Println("查找次数：", i)
		mid := (left + right) / 2
		if mid < 0 || mid >= len(arr) {
			return -1
		}
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			//  找到第一个i
			if mid == 0 || arr[mid-1] != target {
				index = mid
				break
			} else {
				// 递归继续查找
				right = mid - 1
			}
		}
	}
	return index
}

// 最后一个i
func bin_search_midB(arr []int, target int) int {
	fmt.Println(arr, target)
	left := 0
	right := len(arr) - 1
	index := -1
	i := 0
	for left <= right {
		i++
		fmt.Println("查找次数：", i)
		mid := (left + right) / 2
		if mid < 0 || mid >= len(arr) {
			return -1
		}
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			//  找到最后一个i
			if mid == len(arr)-1 || arr[mid+1] != target {
				index = mid
				break
			} else {
				// 递归继续查找
				left = mid + 1
			}
		}
	}
	return index
}

// 第一个大于等于i的
func bin_search_midC(arr []int, target int) int {
	fmt.Println(arr, target)
	left := 0
	right := len(arr) - 1
	index := -1
	i := 0
	for left <= right {
		i++
		fmt.Println("查找次数：", i)
		mid := (left + right) / 2
		if mid < 0 || mid >= len(arr) {
			return -1
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			// 临界点判断
			if mid == 0 || arr[mid-1] < target {
				index = mid
				break
			} else {
				right = mid - 1
			}
		}
	}
	return index
}

// 最后一个小于等于i
func bin_search_midD(arr []int, target int) int {
	fmt.Println(arr, target)
	left := 0
	right := len(arr) - 1
	index := -1
	i := 0
	for left <= right {
		i++
		fmt.Println("查找次数：", i)
		mid := (left + right) / 2
		if mid < 0 || mid >= len(arr) {
			return -1
		}
		if arr[mid] > target {
			right = mid - 1
		} else {
			//  找到最后一个i
			if mid == len(arr)-1 || arr[mid+1] > target {
				index = mid
				break
			} else {
				// 递归继续查找
				left = mid + 1
			}
		}
	}
	return index
}