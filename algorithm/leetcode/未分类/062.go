/*
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
	机器人每次只能向下或者向右移动一步。 机器人试图达到网格的右下角（在下图中标记为“Finish”）。
	问总共有多少条不同的路径？
*/
func uniquePaths(m int, n int) int {
	vector := make([][]int, m)
	for i := 0; i < m; i++ {
		vector[i] = make([]int, n)
	}
	// 初始化第一行第一列
	for row := 0; row < m; row++ {
		vector[row][0] = 1
	}
	for col := 0; col < n; col++ {
		vector[0][col] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			vector[i][j] = vector[i-1][j] + vector[i][j-1]
		}
	}
	return vector[m-1][n-1]
}
