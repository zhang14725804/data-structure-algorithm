package leetcode

// Leetcode074 Search a 2D Matrix（常规方法和二分法）
func Leetcode074() bool {
	matrix := [][]int{
		{1, 3},
	}
	target := 2

	// 常规方法。。将二维数组转化为一维数组
	// var arr []int
	// for i := 0; i < len(matrix); i++ {
	// 	for j := 0; j < len(matrix[i]); j++ {
	// 		arr = append(arr, matrix[i][j])
	// 	}
	// }
	// for i := 0; i < len(arr); i++ {
	// 	if arr[i] == target {
	// 		return true
	// 	}
	// }
	// return false

	/*
		二分法（注意点：临界判断，二维数组转化为一维数组）
		分开判断m和n，因为matrix为[]的情况
	*/
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	left, right := 0, m*n-1
	for left < right {
		mid := (left + right) >> 1
		// 二维数组转化为一维数组
		if matrix[mid/m][mid%m] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if matrix[right/m][right%m] != target {
		return false
	}
	return true
}
