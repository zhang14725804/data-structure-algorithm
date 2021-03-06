/*
	给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
*/
// m、n表示网格的长宽
var n,m int
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0{
		return 0
	}
	n = len(grid)
	m = len(grid[0])
	res:=0
	// 遍历表格
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			if grid[i][j] == '1'{
				res++
				dfs(grid,i,j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte,x,y int){
	// 四个方向
	dx:=[4]int{-1,0,1,0}
	dy:=[4]int{0,1,0,-1}
	// 标记已经走过的位置
	grid[x][y] = '0'
	// 遍历四个方向
	for i:=0;i<4;i++{
		a:=x+dx[i]
		b:=y+dy[i]
		// 边界条件
		if a>=0 && a<n && b>=0 && b<m && grid[a][b] == '1'{
			dfs(grid,a,b)
		}
	}
}