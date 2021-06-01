/*
	方法1：回溯法

	皇后们的约束条件：
	不能同行
	不能同列
	不能同斜线
*/

var ans [][]string
var board [][]byte
var n int

func solveNQueens(_n int) [][]string {
	// 初始化
	n = _n
	ans = make([][]string, 0)
	board = make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	// 从第0行开始
	backtrack(0)
	return ans
}

/*
   row: 是当前递归到棋牌的第几行了
*/
func backtrack(row int) {
	if row == n {
		var str []string
		for i := 0; i < n; i++ {
			str = append(str, string(board[i]))
		}
		ans = append(ans, str)
		return
	}
	// 递归深度就是row控制棋盘的行，每一层里for循环的col控制棋盘的列，一行一列，确定了放置皇后的位置。每次都是要从新的一行的起始位置开始搜，所以都是从0开始。
	for col := 0; col < n; col++ {
		if valid(row, col) {
			board[row][col] = 'Q'
			backtrack(row + 1)
			board[row][col] = '.'
		}
	}
}

/*
	验证棋牌是否合法😅😅😅
	为什么没有在同行进行检查呢：因为在单层搜索的过程中，每一层递归，只会选for循环（也就是同一行）里的一个元素，所以不用去重了
*/
func valid(row, col int) bool {
	// 检查列
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查 45度角是否有皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 检查 135度角是否有皇后
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}