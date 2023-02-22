/*
	方法1：79题（从 words 中依次选定一个单词 -> 从图中的每个位置出发，看能否找到这个单词），再加一个循环

	todo：方法2：从图中的每个位置出发 -> 看遍历过程中是否遇到了 words 中的某个单词。用到了前缀树😅
*/
var board [][]byte
var n, m int
var st [][]bool              // 标记是否已经走过
var dx = [4]int{-1, 0, 1, 0} // 上下左右四个方向
var dy = [4]int{0, 1, 0, -1}

func findWords(_board [][]byte, words []string) []string {
	board = _board
	n, m = len(board), len(board[0])
	res := make([]string, 0)
	for i := 0; i < len(words); i++ {
		if exist(words[i]) {
			res = append([]string{words[i]}, res...)
		}
	}
	return res
}

func exist(word string) bool {
	// 初始化
	st = make([][]bool, n)
	for i := 0; i < n; i++ {
		st[i] = make([]bool, m)
	}
	// 从图中的每个位置出发，看能否找到这个单词
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				if dfs(word, i, j, 1) == true {
					return true
				}
			}
		}
	}
	return false
}

func dfs(word string, i, j, u int) bool {
	if u == len(word) {
		return true
	}
	st[i][j] = true
	for k := 0; k < 4; k++ {
		x := i + dx[k]
		y := j + dy[k]
		if x >= 0 && x < n && y >= 0 && y < m && st[x][y] == false && board[x][y] == word[u] {
			if dfs(word, x, y, u+1) == true {
				return true
			}
		}
	}
	st[i][j] = false
	return false
}