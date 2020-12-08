/*
	1391，判断走通的条件目前想不出来
*/
func hasValidPath(grid [][]int) bool {
    return dfs(0,0,grid)
}
// x，y方向
var dx= [4]int{-1,0,1,0}
var dy= [4]int{0,1,0,-1}

var state= [6][4]int{
	{0,1,0,1},
	{1,0,1,0},
	{0,0,1,1},
	{0,1,1,0},
	{1,0,0,1},
	{1,1,0,0},
}
func dfs(x,y int,grid [][]int) bool{
	// 长宽
	n:=len(grid)
	m:=len(grid[0])
	// 走到右下角
	if x==n-1 && y == m-1{
		return true
	}

	k:=grid[x][y]
	// 标记走过的
	grid[x][y] = 0
	// 遍历四个方向
	for i:=0;i<4;i++{
		a:=x+dx[i]
		b:=y+dy[i]
		// 判断越界或者已经走过
		if a<0 || a>=n || b<0 || b>=m || grid[a][b] == 0{
			continue
		}
		/*
			可以走通的条件，todo：反方向（异或）
			0，1，2，3四个方向
		*/ 
		if state[k-1][i]==0 || state[grid[a][b]-1][i^2] == 0 {
			continue
		}
		if dfs(a,b,grid){
			return true
		}
	}
	return false
}