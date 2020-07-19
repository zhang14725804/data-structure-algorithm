/*
	编写一个程序，通过已填充的空格来解决数独问题。

	一个数独的解法需遵循如下规则：

	数字 1-9 在每一行只能出现一次。
	数字 1-9 在每一列只能出现一次。
	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
	空白格用 '.' 表示。

	给定的数独序列只包含数字 1-9 和字符 '.' 。
	你可以假设给定的数独只有唯一解。
	给定数独永远是 9x9 形式的。
*/
var row [9][9]bool // 每行
var col [9][9]bool // 每列
var cell [3][3][9]bool // 每个小九宫格3*3
func solveSudoku(board [][]byte)  {
	// 遍历当前存在的数
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j]!= '.'{
				t:=board[i][j] - '1'
				// 标记已经存在
				row[i][t] , col[j][t] , cell[i/3][j/3][t] = true,true,true
			}
		}
	}
	// 左上角开始遍历
	dfs(board,0,0)
}

func dfs(board [][]byte,x,y int) bool {
	// 移动纵坐标，判断是否越界
	if y == 9{
		x++
		y = 0
	}
	// 遍历完所有行
	if x == 9{
		return true
	}
	// 已经填了数字，调到下一个位置
	if board[x][y] != '.'{
		return dfs(board,x,y+1)
	}
	// 当前这个位置没有填数字，枚举当前位置可以填那个数字
	for i := 0; i < 9; i++ {
		// 保证每行每列每个九宫格没有出现过这个数字
		if row[x][i]==false && col[y][i]==false && cell[x/3][y/3][i] ==false{
			// 填数字
			board[x][y]=byte('1'+i)
			row[x][i] , col[y][i] , cell[x/3][y/3][i] = true,true,true
			// 下一个格子
			if dfs(board,x,y+1) == true{
				return true
			}
			// 恢复现场
			board[x][y]='.'
			row[x][i] , col[y][i] , cell[x/3][y/3][i] = false,false,false
		}
	}
	return false
}