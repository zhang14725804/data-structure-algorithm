/*
	二分法+双指针
	1. 二分法找到right
	2. 双指针，x-arr[left] <= arr[right]-x left减1，否则right减1
*/
func findClosestElements(arr []int, k, x int) []int {
	// 二分法
	right := sort.SearchInts(arr, x)
	left := right - 1
	for ; k > 0; k-- {
		if left < 0 {
			right++
		} else if right >= len(arr) || x-arr[left] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}
	// 😅 注意返回区间
	return arr[left+1 : right]
}