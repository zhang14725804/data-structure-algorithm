/*
	给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
	找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

	todo：经典的flood fill 问题！
*/
var n,m int
var board [][]byte
var dx = [4]int{-1,0,1,0}
var dy = [4]int{0,1,0,-1}

func solve(_board [][]byte) {
	board = _board
	n = len(board)
	if n ==0 {
		return 
	}
	m = len(board[0])

	// 左右两边
	for i:=0;i<n;i++{
		if board[i][0] == 'O'{
			dfs(i,0)
		}
		if board[i][m-1] == 'O'{
			dfs(i,m-1)
		}
	}

	// 上下两边
	for i:=0;i<m;i++{
		if board[0][i] == 'O'{
			dfs(0,i)
		}
		if board[n-1][i] == 'O'{
			dfs(n-1,i)
		}
	}

	// 中间的内容
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			if board[i][j] == '#'{
				board[i][j] = 'O'
			}else{
				board[i][j] = 'X'
			}
		}
	}

	_board = board
}

func dfs(x int, y int ){
	board[x][y] = '#'
	for i:=0;i<4;i++{
		a := x+dx[i]
		b := y+dy[i]
		if a>=0 && a<n && b>=0 && b<m && board[a][b] == 'O'{
			dfs(a,b)
		} 
	}
}

