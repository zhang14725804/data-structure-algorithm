/*
	😅😅😅 找规律
*/
func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	i := 0 // x+y
	res := []int{}
	for i < m+n {
		// 奇数趟
		// x，y起始值
		x1 := 0
		if i < m {
			x1 = i
		} else {
			x1 = m - 1
		}
		y1 := i - x1
		for x1 >= 0 && y1 < n {
			res = append(res, mat[x1][y1])
			x1--
			y1++
		}
		i++
		if i >= m+n {
			break
		}

		// 偶数趟
		// x，y起始值
		y2 := 0
		if i < n {
			y2 = i
		} else {
			y2 = n - 1
		}
		x2 := i - y2
		for y2 >= 0 && x2 < m {
			res = append(res, mat[x2][y2])
			x2++
			y2--
		}
		i++
	}
	return res
}