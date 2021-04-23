/*
	编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
	每行中的整数从左到右按升序排列。
	每行的第一个整数大于前一行的最后一个整数。
*/

// 精确查找 所以采用【模板1】
func searchMatrix(matrix [][]int, target int) bool {
	// n、m分别表示行、列
	n, m := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return false
	}
	// left、right取值范围
	left, right := 0, n*m-1
	for left < right {
		mid := (left + right) >> 1
		// 😅😅😅😅😅😅 question 如何把编号变成坐标,用以获取元素；mid/m是行，mid%m是列
		// 模板1；精确查找。[l, r]区间划分为[l, mid] 和 [mid+1, r]
		if matrix[mid/m][mid%m] >= target {
			right = mid // 满足条件，移动right
		} else {
			left = mid + 1 // 不满足条件，移动left
		}
	}
	return matrix[left/m][left%m] == target
}