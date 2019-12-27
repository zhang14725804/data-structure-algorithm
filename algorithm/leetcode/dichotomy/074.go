package leetcode

// Leetcode074 Search a 2D Matrix（常规方法和二分法）
func Leetcode074() bool {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	target := 6

	// 常规方法。。将二维数组转化为一维数组
	var arr []int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			arr = append(arr, matrix[i][j])
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}
