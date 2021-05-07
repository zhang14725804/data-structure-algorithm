/*
	思路：递归 question  😅😅😅
	求第 0 层到第 n 层的和最小，
	就是第0层的数字加上第1层到第n层的的最小和。
	递归出口就是，第n层到第n层最小的和，就是该层的数字本身。
*/
var triangle [][]int

func minimumTotal(_triangle [][]int) int {
	triangle = _triangle
	// 从 [0,0] 开始递归
	return dfs(0, 0)
}

func dfs(row, col int) int {
	// base case 😅
	if row == len(triangle) {
		return 0
	}
	min := (1 << 32)
	// 当前行
	curRow := triangle[row]
	// 向下走
	min = MinInt(min, curRow[col]+dfs(row+1, col))
	// 向右下方走
	if col+1 < len(curRow) {
		min = MinInt(min, curRow[col+1]+dfs(row+1, col+1))
	}
	return min
}