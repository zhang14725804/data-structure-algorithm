/*
	岛屿问题
*/
var m,n int
func countBattleships(board [][]byte) int {
	if len(board) == 0 || len(board[0])==0{
		return 0
	}
	m=len(board)
	n=len(board[0])
	res:=0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X'{
				res++
                dfs(board,i,j)
			}
		}
	}
	return res
}
func dfs(board [][]byte,x,y int) {
	dx:=[]int{-1,0,1,0}
	dy:=[]int{0,1,0,-1}
	board[x][y] = '.'
	for i := 0; i < 4; i++ {
		a:=x+dx[i]
		b:=y+dy[i]
		if a>=0 && a<m && b>=0 && b<n && board[a][b] == 'X'{
			dfs(board,a,b)
		}
	}
}