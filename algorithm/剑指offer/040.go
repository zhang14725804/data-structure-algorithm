/*
	蛇形，顺时针打印矩阵
	todo：经典矩阵问题
*/
func printMatrix(matrix [][]int) []int {
	res:=make([]int,0)
	n := len(matrix)
	if n ==0 {
		return res
	}
	m := len(matrix[0])

	// 二维数组，保存已经是否走过的路径
	st:=make([][]bool,n)
	for i:=0; i<n; i++  {
		st[i] = make([]bool,m)
	}
	// 上右下左四个方向
	dx,dy:=[4]int{-1,0,1,0},[4]int{0,1,0,-1}
	// d 初始方向
	x,y,d := 0,0,1

	for i:=0; i<n*m; i++{
		res = append(res,matrix[x][y])
		st[x][y] = true // 标记已经走过
		a,b := x+dx[d],y+dy[d]
		// 判断边界（左右超出边界、已经走过）
		if a<0 || a>=n || b<0 || b>=m || st[a][b]{
			d = (d+1) % 4 // 下一行
			a = x + dx[d]
			b = y + dy[d]
		}
		x = a
		y = b
	}
	return res
}