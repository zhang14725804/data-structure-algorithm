/*
	给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
*/
func generateMatrix(n int) [][]int {
	// 左右上下
	l, r, t, b := 0, n-1, 0, n-1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	num, end := 1, n*n
	for num <= end {
		// 从左到右
		for i := l; i <= r; i++ {
			res[t][i] = num
			num++
		}
		t++
		// 从上到下
		for i := t; i <= b; i++ {
			res[i][r] = num
			num++
		}
		r--
		// 从右到左
		for i := r; i >= l; i-- {
			res[b][i] = num
			num++
		}
		b--
		// 从下到上
		for i := b; i >= t; i-- {
			res[i][l] = num
			num++
		}
		l++
	}
	return res
}