package common

/*
	常用到的一些东西
*/
func compare(a, b int, max bool) int {
	// max 是否返回最大值
	if a > b {
		if max == true {
			return a
		}
		return b
	}
	if max == true {
		return b
	}
	return a
}

// int最大值，最小值 -9223372036854775808 9223372036854775807
const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
