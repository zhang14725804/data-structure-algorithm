/*
	todo：不怎么懂
*/
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
	// 几个意思
	grid[x][y] = '0'
	for i:=0;i<4;i++{
		a:=x+dx[i]
		b:=y+dy[i]
		if a>=0 && a<n && b>=0 && b<m && grid[a][b] == '1'{
			dfs(grid,a,b)
		}
	}
}