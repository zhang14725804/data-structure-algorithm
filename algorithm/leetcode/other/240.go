/*
	编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

	每行的元素从左到右升序排列。
	每列的元素从上到下升序排列。

*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	for i := 0; i < n; i++ {
		// 提前跳出循环
		if matrix[i][0] > target {
			break
		}
		if matrix[i][m-1] < target {
			continue
		}
		col := binarySearch(matrix[i], target)
		//fmt.Println(i,col,matrix[i][col])
		if matrix[i][col] == target {
			return true
		}
	}
	return false
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		// 是否存在一个目标值（精确查找）
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}