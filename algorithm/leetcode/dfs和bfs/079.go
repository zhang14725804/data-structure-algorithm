/*
	79. Word Search
	不能走重复的路线

	todos：：这种题，我遇到就蒙圈了
	这种声明初始化为什么不行
	var dy []int = {0,1,0,-1}
*/
var m,n int // 矩阵的长宽
var dx []int
var dy []int
func exist(board [][]byte, word string) bool {
	// 上右下左
	dx = []int{-1,0,1, 0}
	// 上右下左
	dy = []int{0,1,0,-1}
	// 边界情况
	if len(board) ==0 || len(board[0])==0{
		return false
	}
	// 行列
	n = len(board)
	m = len(board[0])

	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			if dfs(board,i,j,word,0) == true {
				return true
			}
		}
	}
	return false
}
/*
	x,y当前走到的位置
	word目标单词
	u目标单词第几位
*/ 
func dfs(board [][]byte,x,y int,word string,u int) bool{
	// 当前单词不匹配
	if board[x][y]!=word[u]{
		return false
	}
	// 单词的最后一个匹配完成
	if u ==len(word) -1{
		return true
	}
	// 回溯（恢复初始状态）,保证每条线路看到的状态都是一样的
	board[x][y] = '.'
	// 枚举上下左右四个方向
	for i:=0;i<4;i++{
		// 当前每个方向的坐标
		a:=x+dx[i]
		b:=y+dy[i]
		// 是否在界内
		if a>=0 && a<n && b>=0&&b<m{
			// 走到下一个格子是否可以匹配
			if dfs(board,a,b,word,u+1) == true {
				return true
			}
		}
	}
	// 
	board[x][y] = word[u]
	return false
}