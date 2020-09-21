/*
	给定一个二维网格和一个单词，找出该单词是否存在于网格中。
	方法：从 words 中依次选定一个单词 -> 从图中的每个位置出发，看能否找到这个单词
*/
var n, m int
var board [][]byte
var st [][]bool //  标记是否已经走过
var word string

func exist(_board [][]byte, _word string) bool {
	board = _board
	word = _word
	n = len(board)
	m = len(board[0])
	// 初始化二维数组
	st = make([][]bool, n)
	for i := 0; i < n; i++ {
		st[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 起始位置是word首字母（都是byte可以直接比较）
			if board[i][j] == word[0] {
				if dfs(i, j, 1) == true {
					return true
				}
			}
		}
	}
	return false
}

func dfs(x, y, u int) bool {
	if u == len(word) {
		return true
	}
	// 标记已经走过
	st[x][y] = true
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}
	// 遍历四个方向
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		// 判断临界点
		if a >= 0 && a < n && b >= 0 && b < m && st[a][b] == false && board[a][b] == word[u] {
			if dfs(a, b, u+1) == true {
				return true
			}
		}
	}
	// 恢复初始状态
	st[x][y] = false
	return false
}