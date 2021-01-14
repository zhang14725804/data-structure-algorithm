/*
	给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

	84，85，221 难以理解😅

*/

// 方法1：暴力破解
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 { // []的情况
		return 0
	}
	n, m := len(matrix), len(matrix[0])
	// 保存以当前数字结尾的连续 1 的个数
	width := make([][]int, n)
	for i := 0; i < n; i++ {
		width[i] = make([]int, m)
	}

	area := 0
	// 遍历每个元素
	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if matrix[row][col] == '1' {
				if col == 0 {
					width[row][col] = 1
				} else {
					width[row][col] = width[row][col-1] + 1
				}
			} else {
				width[row][col] = 0
			}

			// 记录所有行中最小的数
			minWidth := width[row][col]
			// 向上扩展
			for up := row; up >= 0; up-- {
				height := row - up + 1
				// 找最小的数作为矩阵的宽
				minWidth = MinInt(minWidth, width[up][col])
				area = MaxInt(area, height*minWidth)
			}
		}
	}
	return area
}

// 方法2：参考84题
