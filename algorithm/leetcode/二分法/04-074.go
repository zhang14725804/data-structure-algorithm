/**
	精确查找
	
	编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
	每行中的整数从左到右按升序排列。
	每行的第一个整数大于前一行的最后一个整数。

*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0{
		return false
	}
	n,m := len(matrix),len(matrix[0])
	l,r := 0, n*m - 1
	// 技巧：如何把编号变成坐标,用以获取元素
	for l<r{
		mid := (l+r) >> 1
		// 模板1
		if matrix[mid / m][mid % m] >= target{
			r = mid
		}else{
			l = mid+1
		}
	}
	return matrix[r / m][r % m] == target
}