/*
	给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

	思路：右下左上循环遍历
*/
func spiralOrder(m [][]int) []int {
	top:=0
	bottom:=len(m)-1
	left:=0
	right:=len(m[0])-1

	// 遍历方向 right，bottom，left，top
	dir:="right"
	// 遍历就好了
	for left<=rihgt && top<=bottom{
		
	}
}
