/*
	给定一个 n × n 的二维矩阵表示一个图像。
	将图像顺时针旋转 90 度。

	思路：45度交换，然后水平交换。
*/
func rotate(matrix [][]int)  {
	// 45度交换
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 水平方向
	for i := 0; i < len(matrix); i++ {
		for j,k := 0,len(matrix[i])-1; j < k; j,k = j+1,k-1{
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}
